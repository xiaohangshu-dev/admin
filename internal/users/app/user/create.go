package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	"github.com/xiaohangshuhub/admin/internal/users/domain/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CreateCmd 创建用户命令,包含创建用户所需的信息
type CreateCmd struct {
	Username string      // 用户名
	Nickname string      // 昵称
	Avatar   string      // 头像
	Email    *string     // 邮箱
	Phone    *string     // 手机号
	Pwd      string      // 密码
	Salt     string      // 密码盐值
	Gender   user.Gender // 性别
	Roles    []uuid.UUID // 角色
}

type CreateCmdHandler struct {
	um *user.Manager
	pm *roleperm.Manager
	*gorm.DB
	*zap.Logger
}

// NewCreateCmdHandler 返回创建用户命令处理器
func NewCreateCmdHandler(um *user.Manager, pm *roleperm.Manager, db *gorm.DB, zap *zap.Logger) *CreateCmdHandler {
	return &CreateCmdHandler{
		um:     um,
		pm:     pm,
		DB:     db,
		Logger: zap,
	}
}

// Handle 处理创建用户命令
func (c *CreateCmdHandler) Handle(ctx context.Context, cmd CreateCmd) (bool, error) {

	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return false, errors.New("invalid user id in context")
	}

	u, err := c.um.Create(cmd.Username, cmd.Nickname, cmd.Avatar, cmd.Pwd, cmd.Phone, cmd.Email, uid, cmd.Gender)

	if err != nil {
		return false, err
	}

	if result := c.DB.Create(u); result.Error == nil {
		return false, err
	}

	if err := c.pm.ConfigureUserRoles(u.ID, cmd.Roles); err != nil {
		return false, err
	}

	return true, nil
}
