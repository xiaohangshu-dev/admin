package roleperm

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshu-dev/admin/internal/users/domain/roleperm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// RoleCreateCmd 创建角色命令
type RoleCreateCmd struct {
	ParentID uuid.UUID   // 父节点ID
	Role     string      // 角色
	Name     string      // 名
	Desc     string      // 头像
	PermIds  []uuid.UUID // 权限
}

type RoleCreateCmdHandler struct {
	*roleperm.Manager
	*gorm.DB
	*zap.Logger
}

// NewRoleCreateCmdHandler 返回创建角色命令处理器
func NewRoleCreateCmdHandler(manager *roleperm.Manager, db *gorm.DB, zap *zap.Logger) *RoleCreateCmdHandler {
	return &RoleCreateCmdHandler{
		Manager: manager,
		DB:      db,
		Logger:  zap,
	}
}

// Handle 处理创建角色命令
func (c *RoleCreateCmdHandler) Handle(ctx context.Context, cmd RoleCreateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(string)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	r, err := c.Manager.CreateRole(cmd.Role, cmd.Name, uid, cmd.ParentID)

	if err != nil {
		return false, err
	}

	rp, err := c.Manager.ConfigureRolePerms(r.ID, cmd.PermIds)

	if err != nil {
		return false, err
	}

	// 开启事务
	if err := c.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(r).Error; err != nil {
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
