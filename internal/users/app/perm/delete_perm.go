package perm

import (
	"context"

	"github.com/xiaohangshuhub/admin/internal/users/domain/perm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeleteCmd 删除用户命令,包含删除用户所需的信息
type PermDeleteCmd struct {
	ID string `json:"id"`
}

type PermDeleteCmdHandler struct {
	*gorm.DB
	*zap.Logger
}

func NewPermDeleteCmdHandler(repo *gorm.DB, zap *zap.Logger) *PermDeleteCmdHandler {
	return &PermDeleteCmdHandler{
		DB:     repo,
		Logger: zap,
	}
}

func (h *PermDeleteCmdHandler) Handle(ctx context.Context, cmd PermDeleteCmd) (bool, error) {

	if tx := h.DB.Delete(&perm.Permission{}, cmd.ID); tx.Error != nil {
		h.Logger.Error("db delete user failed", zap.String("ID", cmd.ID), zap.Error(tx.Error))
		return false, tx.Error
	}

	return true, nil
}
