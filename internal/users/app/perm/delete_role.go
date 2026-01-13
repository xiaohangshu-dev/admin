package perm

import (
	"context"

	"github.com/xiaohangshuhub/admin/internal/users/domain/perm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeleteCmd 删除用户命令,包含删除用户所需的信息
type RoleDeleteCmd struct {
	ID string `json:"id"`
}

type RoleDeleteCmdHandler struct {
	*gorm.DB
	*zap.Logger
}

func NewRoleDeleteCmdHandler(repo *gorm.DB, zap *zap.Logger) *RoleDeleteCmdHandler {
	return &RoleDeleteCmdHandler{
		DB:     repo,
		Logger: zap,
	}
}

func (h *RoleDeleteCmdHandler) Handle(ctx context.Context, cmd RoleDeleteCmd) (bool, error) {

	if tx := h.DB.Delete(&perm.Role{}, cmd.ID); tx.Error != nil {
		h.Logger.Error("db delete user failed", zap.String("ID", cmd.ID), zap.Error(tx.Error))
		return false, tx.Error
	}

	return true, nil
}
