package user

import (
	"context"

	"github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeleteCmd 删除用户命令,包含删除用户所需的信息
type DeleteCmd struct {
	ID string `json:"id"`
}

type DeleteCmdHandler struct {
	*gorm.DB
	*zap.Logger
}

func NewDeleteCmdHandler(repo *gorm.DB, zap *zap.Logger) *DeleteCmdHandler {
	return &DeleteCmdHandler{
		DB:     repo,
		Logger: zap,
	}
}

func (h *DeleteCmdHandler) Handle(ctx context.Context, cmd DeleteCmd) (bool, error) {
	tx := h.DB.Delete(&user.Account{}, cmd.ID)

	if tx.Error != nil {
		h.Logger.Error("db delete user failed", zap.String("ID", cmd.ID), zap.Error(tx.Error))
		// TODO: 后期优化
		return false, tx.Error
	}

	return true, nil
}
