package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/gender"
	"github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UpdateCmd 更新用户命令,包含更新用户所需的信息
type UpdateCmd struct {
	ID       uuid.UUID     `json:"id"`
	Nikename string        `json:"nikename"`
	Roles    []uuid.UUID   `json:"roles"`
	Avatar   string        `json:"avatar"`
	Email    *string       `json:"email"`
	Phone    *string       `json:"phone"`
	Pwd      string        `json:"pwd"`
	Salt     string        `json:"salt"`
	Gender   gender.Gender `json:"gender"`
}

type UpdateCmdHandler struct {
	*user.Manager
	*gorm.DB
	*zap.Logger
}

func NewUpdateCmdHandler(m *user.Manager, repo *gorm.DB, zap *zap.Logger) *UpdateCmdHandler {
	return &UpdateCmdHandler{
		Manager: m,
		DB:      repo,
		Logger:  zap,
	}
}

func (h *UpdateCmdHandler) Handle(ctx context.Context, cmd UpdateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	account, err := h.Manager.Update(cmd.ID, cmd.Nikename, cmd.Avatar, cmd.Phone, cmd.Email, uid, cmd.Gender, cmd.Roles)

	if err != nil {
		return false, err
	}

	tx := h.DB.Save(account)

	if tx.Error != nil {
		// TODO: 后续优化
		return false, tx.Error
	}

	return true, nil
}
