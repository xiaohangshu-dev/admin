package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	"github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"gorm.io/gorm"
)

type UserInfoDto struct {
	ID       uuid.UUID
	Username string
	Email    *string
	Phone    *string
	Name     string
	Avatar   string
	roles    []uuid.UUID
}

type UserQuery struct {
}

type UserQueryHandler struct {
	DB *gorm.DB
}

func NewUserQueryHandler(repo *gorm.DB) *UserQueryHandler {
	return &UserQueryHandler{
		DB: repo,
	}
}

func (h *UserQueryHandler) Handle(ctx context.Context, query UserQuery) (UserInfoDto, error) {
	uid, ok := ctx.Value("UserID").(uuid.UUID)

	if !ok {
		return UserInfoDto{}, errors.New("invalid user id in context")
	}

	var user user.Account

	if tx := h.DB.First(&user, uid); tx.Error != nil {
		return UserInfoDto{}, tx.Error
	}

	var roles []roleperm.UserRole

	if tx := h.DB.Where("user_id = ?", user.ID).Find(&roles); tx.Error != nil {
		return UserInfoDto{}, tx.Error
	}

	var roleIds []uuid.UUID

	for _, role := range roles {
		roleIds = append(roleIds, role.RoleID)
	}

	return UserInfoDto{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Phone:    user.Phone,
		roles:    roleIds,
	}, nil

}
