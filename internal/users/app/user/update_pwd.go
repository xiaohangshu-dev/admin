package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshu-dev/admin/internal/users/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UpdateCmd 更新用户命令,包含更新用户所需的信息
type UpdatePwdCmd struct {
	ID  uuid.UUID `json:"id"`
	Pwd string    `json:"pwd"`
}

type UpdatePwdCmdHandler struct {
	*user.Manager
	*gorm.DB
	*zap.Logger
}

func NewUpdatePwdCmdHandler(m *user.Manager, repo *gorm.DB, zap *zap.Logger) *UpdatePwdCmdHandler {
	return &UpdatePwdCmdHandler{
		Manager: m,
		DB:      repo,
		Logger:  zap,
	}
}

func (h *UpdatePwdCmdHandler) Handle(ctx context.Context, cmd UpdatePwdCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	account, err := h.Manager.UpdatePwd(cmd.ID, cmd.Pwd, uid)

	if err != nil {
		return false, err
	}

	if tx := h.Save(account); tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
