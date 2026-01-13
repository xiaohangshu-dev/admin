package perm

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/perm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PermCreateCmd 创建权限命令,包含创建权限所需的信息
type PermCreateCmd struct {
	Name     string     // 用户名
	Route    string     // 昵称
	Icon     string     // 图标
	Desc     string     // 头像
	Weight   int32      // 权重
	Type     perm.Type  // 功能类型
	ParentID *uuid.UUID // 父节点ID
}

type PermCreateCmdHandler struct {
	*perm.Manager
	*gorm.DB
	*zap.Logger
}

// NewPermCreateCmdHandler 返回创建权限命令处理器
func NewPermCreateCmdHandler(manager *perm.Manager, db *gorm.DB, zap *zap.Logger) *PermCreateCmdHandler {
	return &PermCreateCmdHandler{
		Manager: manager,
		DB:      db,
		Logger:  zap,
	}
}

// Handle 处理创建权限命令
func (c *PermCreateCmdHandler) Handle(ctx context.Context, cmd PermCreateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}
	u, err := c.Manager.CreatePremission(cmd.Name, cmd.Route, cmd.Icon, cmd.Desc, cmd.Weight, cmd.Type, cmd.ParentID, uid)

	if err != nil {
		return false, err
	}

	if err := c.DB.Create(u).Error; err != nil {
		return false, err
	}

	return true, nil
}
