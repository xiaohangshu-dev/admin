package user

import (
	"context"
	"errors"

	"github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Login struct {
	Username string
	Password string
}

type LoginHandler struct {
	*gorm.DB
	*zap.Logger
}

func NewLoginHandler(repo *gorm.DB, zap *zap.Logger) *LoginHandler {

	return &LoginHandler{
		DB:     repo.Model(&user.Account{}),
		Logger: zap,
	}

}

// Handler  根据登录名和密码查询用户,返回用户数据传输对象或错误信息。
func (h *LoginHandler) Handle(ctx context.Context, req Login) (UserDto, error) {

	var acc user.Account

	err := h.Where("username = ? OR phone = ? OR email = ?", req.Username, req.Username, req.Username).First(&acc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserDto{}, user.ErrUserNotFound
		}
		return UserDto{}, err
	}

	if !acc.CheckPassword(req.Password) {
		return UserDto{}, user.ErrPasswordInvalid
	}

	return UserDto{
		UserName:    acc.Username,
		Nikename:    acc.Nickname,
		Email:       acc.Email,
		Phone:       acc.Phone,
		AccessToken: "123", // TODO: 这里后续换成 JWT
	}, nil
}
