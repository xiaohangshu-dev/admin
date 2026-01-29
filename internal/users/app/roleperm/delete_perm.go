package roleperm

import (
	"context"

	"github.com/google/uuid"
	"github.com/xiaohangshu-dev/admin/internal/users/domain/roleperm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DeleteCmd 删除用户命令,包含删除用户所需的信息
type PermDeleteCmd struct {
	ID uuid.UUID `json:"id"`
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

	if tx := h.Delete(&roleperm.Permission{}, cmd.ID); tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
