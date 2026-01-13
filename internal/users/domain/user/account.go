package user

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/status"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

type Gender int8

const (
	Unknown Gender = iota + 1 // 保密
	Male                      // 男
	Female                    // 女
)

// Account 描述用户账户领域对象
type Account struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Username                     string        // 用户名
	Nickname                     string        // 昵称
	Avatar                       string        // 头像
	Email                        *string       // 邮箱
	Phone                        *string       // 手机号
	Pwd                          string        // 密码
	Salt                         string        // 密码盐值
	Gender                       Gender        // 性别
	Status                       status.Status // 状态
	CreateBy                     uuid.UUID     // 创建人
	UpdateBy                     *uuid.UUID    // 更新人
}

// newAccount 创建账户并返回实例.
func newAccount(username, nickname, avatar, pwd string, createBy uuid.UUID, gender Gender) (*Account, *Error) {

	account := &Account{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Status:        status.Enable,
		Gender:        gender,
	}
	// 业务规则校验
	if username == "" {
		return nil, ErrUsernameEmpty
	}
	if err := account.SetPassword(pwd); err != nil {
		return account, nil
	}
	if err := account.SetNickname(nickname); err != nil {
		return account, err
	}
	if err := account.SetAvatar(avatar); err != nil {
		return account, err
	}
	if createBy == uuid.Nil {
		return nil, ErrCreateByEmpty
	}
	// 赋值
	account.Username = username
	account.CreateBy = createBy
	return account, nil
}

// SetNickname 设置昵称
func (a *Account) SetNickname(nickname string) *Error {
	if nickname == "" {
		return ErrPwdEmpty
	}

	a.Nickname = nickname
	return nil
}

// SetNickname 设置昵称
func (a *Account) SetAvatar(avatar string) *Error {
	if avatar == "" {
		return ErrAvatarEmpty
	}
	a.Avatar = avatar

	return nil
}

// SetPassword 设置密码
func (a *Account) SetPassword(pwd string) *Error {
	if pwd == "" {
		return ErrPwdEmpty
	}

	a.Pwd = pwd
	return nil
}

// DisableAccount 禁用账户
func (a *Account) DisableAccount() {
	a.Status = status.Disable
}

// EnableAccount 启用账户
func (a *Account) EnableAccount() {
	a.Status = status.Enable
}

// IsEnabled 检查账户是否启用
func (a *Account) IsEnabled() bool {
	return a.Status == status.Enable
}

// CheckPassword 检查密码是否正确
func (a *Account) CheckPassword(pwd string) bool {
	// TODO: 密码加密对比
	return a.Pwd == pwd
}
