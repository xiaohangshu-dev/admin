package perm

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/perm"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PermUpdateCmd 更新权限命令
type PermUpdateCmd struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     // 用户名
	Route    string     // 昵称
	Icon     string     // 图标
	Desc     string     // 头像
	Weight   int32      // 权重
	Type     perm.Type  // 功能类型
	ParentID *uuid.UUID // 父节点ID
}

type PermUpdateCmdHandler struct {
	*perm.Manager
	*gorm.DB
	*zap.Logger
}

// NewPermUpdateCmdHandler 返回更新权限命令处理器
func NewPermUpdateCmdHandler(pm *perm.Manager, repo *gorm.DB, zap *zap.Logger) *PermUpdateCmdHandler {
	return &PermUpdateCmdHandler{
		Manager: pm,
		DB:      repo,
		Logger:  zap,
	}
}

// Handle 处理更新权限命令
func (h *PermUpdateCmdHandler) Handle(ctx context.Context, cmd PermUpdateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	perm, err := h.Manager.UpdatePermission(cmd.ID, cmd.Name, cmd.Route, cmd.Icon, cmd.Desc, cmd.Weight, cmd.Type, cmd.ParentID, uid)

	if err != nil {
		return false, err
	}

	if tx := h.DB.Save(perm); tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
