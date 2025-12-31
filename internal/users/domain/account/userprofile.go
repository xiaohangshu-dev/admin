package account

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

type Gender int

const (
	Unknown Gender = iota // 保密
	Male                  // 男
	Female                // 女
)

type UserProfile struct {
	ddd.Entity[uuid.UUID]        // 实体
	Nickname              string // 昵称
	AvatarURL             string // 头像地址
	Gender                Gender // 性别
}

func (u *UserProfile) NewUserProfile(id uuid.UUID, nickname, avatarURL string, gender Gender) *UserProfile {
	return &UserProfile{
		Entity:    ddd.NewEntity(id),
		Nickname:  nickname,
		AvatarURL: avatarURL,
		Gender:    gender,
	}
}
