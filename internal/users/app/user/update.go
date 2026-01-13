package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/perm"
	"github.com/xiaohangshuhub/admin/internal/users/domain/user"
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
	pm *perm.Manager
	*gorm.DB
	*zap.Logger
}

func NewUpdateCmdHandler(um *user.Manager, pm *perm.Manager, repo *gorm.DB, zap *zap.Logger) *UpdateCmdHandler {
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

	if tx := h.DB.Save(account); tx.Error != nil {
		return false, tx.Error
	}

	if err := h.pm.ConfigureUserRoles(account.ID, cmd.Roles); err != nil {
		return false, err
	}

	return true, nil
}
