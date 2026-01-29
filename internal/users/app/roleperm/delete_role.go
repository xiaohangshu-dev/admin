package roleperm

import (
	"context"

	"github.com/google/uuid"
	"github.com/xiaohangshu-dev/admin/internal/users/domain/roleperm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeleteCmd 删除用户命令,包含删除用户所需的信息
type RoleDeleteCmd struct {
	ID uuid.UUID `json:"id"`
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

	if tx := h.Delete(&roleperm.Role{}, cmd.ID); tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
