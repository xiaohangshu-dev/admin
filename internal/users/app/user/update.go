package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshu-dev/admin/internal/users/domain/roleperm"
	"github.com/xiaohangshu-dev/admin/internal/users/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UpdateCmd 更新用户命令,包含更新用户所需的信息
type UpdateCmd struct {
	ID       uuid.UUID   `json:"id"`
	Nikename string      `json:"nikename"`
	Avatar   string      `json:"avatar"`
	Email    *string     `json:"email"`
	Phone    *string     `json:"phone"`
	Pwd      string      `json:"pwd"`
	Salt     string      `json:"salt"`
	Gender   user.Gender `json:"gender"`
	Roles    []uuid.UUID `json:"roles"`
}

type UpdateCmdHandler struct {
	um *user.Manager
	pm *roleperm.Manager
	*gorm.DB
	*zap.Logger
}

func NewUpdateCmdHandler(um *user.Manager, pm *roleperm.Manager, repo *gorm.DB, zap *zap.Logger) *UpdateCmdHandler {
	return &UpdateCmdHandler{
		um:     um,
		pm:     pm,
		DB:     repo,
		Logger: zap,
	}
}

func (h *UpdateCmdHandler) Handle(ctx context.Context, cmd UpdateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	account, err := h.um.Update(cmd.ID, cmd.Nikename, cmd.Avatar, cmd.Phone, cmd.Email, uid, cmd.Gender)

	if err != nil {
		return false, err
	}

	ur, err := h.pm.ConfigureUserRoles(account.ID, cmd.Roles)

	if err != nil {
		return false, err
	}

	// 开启事务
	if err := h.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&roleperm.UserRole{}, "user_id = ?", account.ID).Error; err != nil {
			return err
		}
		if err := tx.Save(account).Error; err != nil {
			return err
		}
		if err := tx.Create(ur).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return false, err
	}

	return true, nil
}
