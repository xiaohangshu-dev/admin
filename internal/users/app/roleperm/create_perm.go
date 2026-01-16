package roleperm

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PermCreateCmd 创建权限命令,包含创建权限所需的信息
type PermCreateCmd struct {
	Title    string        // 用户名
	Perm     string        // 权限
	Route    string        // 昵称
	Icon     string        // 图标
	Desc     string        // 头像
	Weight   int32         // 权重
	Type     roleperm.Type // 功能类型
	ParentID *uuid.UUID    // 父节点ID
}

type PermCreateCmdHandler struct {
	*roleperm.Manager
	*gorm.DB
	*zap.Logger
}

// NewPermCreateCmdHandler 返回创建权限命令处理器
func NewPermCreateCmdHandler(manager *roleperm.Manager, db *gorm.DB, zap *zap.Logger) *PermCreateCmdHandler {
	return &PermCreateCmdHandler{
		Manager: manager,
		DB:      db,
		Logger:  zap,
	}
}

// Handle 处理创建权限命令
func (c *PermCreateCmdHandler) Handle(ctx context.Context, cmd PermCreateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(string)

	if !ok {
		return false, errors.New("invalid user id in context")
	}
	u, err := c.Manager.CreatePremission(cmd.Title, cmd.Route, cmd.Icon, cmd.Desc, uid, cmd.Weight, cmd.Type, cmd.ParentID)

	if err != nil {
		return false, err
	}

	if err := c.Create(u).Error; err != nil {
		return false, err
	}

	return true, nil
}
