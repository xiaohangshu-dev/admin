package user

import (
	"github.com/xiaohangshuhub/xiaohangshu/internal/users/domain/account"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// CreateCmd 创建用户命令,包含创建用户所需的信息
type CreateCmd struct {
	Loginname string  `json:"loginname"`
	Password  string  `json:"password"`
	Createby  string  `json:"createby"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
}

type CreateCmdHandler struct {
	*account.Manager
	*gorm.DB
	*zap.Logger
}

func NewCreateCmdHandler(manager *account.Manager, db *gorm.DB, zap *zap.Logger) *CreateCmdHandler {
	return &CreateCmdHandler{
		Manager: manager,
		DB:      db,
		Logger:  zap,
	}
}

func (c *CreateCmdHandler) Handle(cmd CreateCmd) (bool, error) {

	u, err := c.Manager.Create(*cmd.Phone, cmd.Password, cmd.Createby)

	if err != nil {
		return false, err
	}

	// 保存并记录未知错误
	if err := c.DB.Create(u).Error; err != nil {

		c.Logger.Error("db create user failed", zap.String("loginname", cmd.Loginname), zap.Error(err))

		// 统一返回业务错误
		return false, account.ErrUserCreateFailed
	}

	return true, nil
}
