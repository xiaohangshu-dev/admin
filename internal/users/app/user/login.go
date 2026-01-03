package user

import (
	"context"
	"errors"

	"github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Login struct {
	Phone    string
	Password string
}

type LoginHandler struct {
	*gorm.DB
	*zap.Logger
}

func NewLoginHandler(repo *gorm.DB, zap *zap.Logger) *LoginHandler {

	return &LoginHandler{
		DB:     repo,
		Logger: zap,
	}

}

// Handler  根据登录名和密码查询用户,返回用户数据传输对象或错误信息。
func (h *LoginHandler) Handle(ctx context.Context, req Login) (UserDto, error) {

	acc := &user.Account{}

	result := h.First(acc)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return UserDto{}, user.ErrUserNotFound
	}

	// 记录未知错误
	if result.Error != nil {
		h.Logger.Error("get user by password failed", zap.Error(result.Error))
		return UserDto{}, user.ErrUnknown
	}

	verif := acc.CheckPassword(req.Password)

	if !verif {
		return UserDto{}, user.ErrPasswordInvalid
	}

	return UserDto{}, nil
}
