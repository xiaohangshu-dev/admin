package user

import (
	"context"

	"github.com/xiaohangshu-dev/admin/internal/users/domain/user"
	"gorm.io/gorm"
)

type UserListQuery struct {
	UserName *string
	Page     int
	Size     int
}

type UserListDto struct {
	Total int64
	List  []UserInfoDto
}

type UserListQueryHandler struct {
	*gorm.DB
}

func NewUserListQueryHandler(repo *gorm.DB) *UserListQueryHandler {
	return &UserListQueryHandler{
		DB: repo,
	}
}

func (h *UserListQueryHandler) Handle(ctx context.Context, query UserListQuery) (UserListDto, error) {

	var users []user.Account

	var spec map[string]any = map[string]any{}

	if query.UserName != nil {
		spec["username"] = *query.UserName
	}
	// 查询总数
	var total int64

	if err := h.Model(&user.Account{}).Where(spec).Count(&total).Error; err != nil {
		return UserListDto{}, err
	}

	if err := h.Where(spec).Find(&users).Error; err != nil {
		return UserListDto{}, err
	}

	var list []UserInfoDto

	for _, u := range users {
		list = append(list, UserInfoDto{
			ID:       u.ID,
			Username: u.Username,
			Name:     u.Nickname,
			Avatar:   u.Avatar,
			Email:    u.Email,
			Phone:    u.Phone,
		})
	}

	return UserListDto{
		Total: total,
		List:  list,
	}, nil
}
