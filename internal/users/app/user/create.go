package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CreateCmd 创建用户命令,包含创建用户所需的信息
type CreateCmd struct {
	Username string      // 用户名
	Nickname string      // 昵称
	Roles    []uuid.UUID // 角色
	Avatar   string      // 头像
	Email    *string     // 邮箱
	Phone    *string     // 手机号
	Pwd      string      // 密码
	Salt     string      // 密码盐值
	Gender   user.Gender // 性别
}

type CreateCmdHandler struct {
	*user.Manager
	*gorm.DB
	*zap.Logger
}

func NewCreateCmdHandler(manager *user.Manager, db *gorm.DB, zap *zap.Logger) *CreateCmdHandler {
	return &CreateCmdHandler{
		Manager: manager,
		DB:      db,
		Logger:  zap,
	}
}

func (c *CreateCmdHandler) Handle(ctx context.Context, cmd CreateCmd) (bool, error) {

	// uid, ok := ctx.Value("UserID").(uuid.UUID)

	// if !ok {
	// 	return false, errors.New("invalid user id in context")
	// }

	uid, _ := uuid.Parse("198b6d03-6143-4fb7-866e-2dffaad5affa")

	u, err := c.Manager.Create(cmd.Username, cmd.Nickname, cmd.Avatar, cmd.Pwd, cmd.Phone, cmd.Email, uid, cmd.Gender, cmd.Roles)

	if err != nil {
		return false, err
	}

	if err := c.DB.Create(u).Error; err != nil {
		return false, err
	}

	return true, nil
}
