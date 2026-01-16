package roleperm

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// RoleUpdateCmd 更新角色命令
type RoleUpdateCmd struct {
	ID       uuid.UUID   `json:"id"` // ID
	ParentID uuid.UUID   // 父节点ID
	Role     string      // 角色
	Name     string      // 名称
	Desc     string      // 头像
	PermIds  []uuid.UUID // 权限
}

type RoleUpdateCmdHandler struct {
	*roleperm.Manager
	*gorm.DB
	*zap.Logger
}

// NewRoleUpdateCmdHandler 返回更新角色命令处理器
func NewRoleUpdateCmdHandler(pm *roleperm.Manager, repo *gorm.DB, zap *zap.Logger) *RoleUpdateCmdHandler {
	return &RoleUpdateCmdHandler{
		Manager: pm,
		DB:      repo,
		Logger:  zap,
	}
}

// Handle 处理更新角色命令
func (h *RoleUpdateCmdHandler) Handle(ctx context.Context, cmd RoleUpdateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(string)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	role, err := h.Manager.UpdateRole(cmd.ID, cmd.Role, cmd.Name, uid, cmd.ParentID)

	if err != nil {
		return false, err
	}

	rp, err := h.Manager.ConfigureRolePerms(role.ID, cmd.PermIds)

	if err != nil {
		return false, err
	}

	// 开启事务
	if err := h.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&roleperm.RolePerm{}, "role_id = ?", role.ID).Error; err != nil {
			return err
		}
		if err := tx.Save(role).Error; err != nil {
			return err
		}
		if err := tx.Create(rp).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return false, err
	}

	return true, nil
}
