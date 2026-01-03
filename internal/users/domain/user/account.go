package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/gender"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/status"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

// Account 描述用户账户领域对象
type Account struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Username                     string        // 用户名
	Nickname                     string        // 昵称
	Roles                        []uuid.UUID   // 角色
	Avatar                       string        // 头像
	Email                        *string       // 邮箱
	Phone                        *string       // 手机号
	Pwd                          string        // 密码
	Salt                         string        // 密码盐值
	Gender                       gender.Gender // 性别
	Status                       status.Status // 状态
	CreatedAt                    time.Time     // 创建时间
	CreateBy                     uuid.UUID     // 创建人
	UpdatedAt                    *time.Time    // 更新时间
	UpdatedBy                    *string       // 更新人
}

// newAccount 创建账户并返回实例.
func newAccount(id uuid.UUID, username, nickname, avatar, pwd string, createBy uuid.UUID, gender gender.Gender, roles []uuid.UUID) (*Account, *Error) {

	// 业务规则校验
	if id == uuid.Nil {
		return nil, ErrIDEmpty
	}
	if username == "" {
		return nil, ErrUsernameEmpty
	}
	if createBy == uuid.Nil {
		return nil, ErrCreateByEmpty
	}

	account := &Account{
		AggregateRoot: ddd.NewAggregateRoot(id),
		Username:      username,
		Status:        status.Active,
		Gender:        gender,
		CreatedAt:     time.Now(),
		CreateBy:      createBy,
	}

	if account, err := account.SetPassword(pwd); err != nil {
		return account, nil
	}

	if account, err := account.SetNickname(nickname); err != nil {
		return account, err
	}

	if account, err := account.SetAvatar(avatar); err != nil {
		return account, err
	}

	if account, err := account.SetRoles(roles); err != nil {
		return account, err
	}

	return account, nil
}

// SetNickname 设置昵称
func (a *Account) SetNickname(nickname string) (*Account, *Error) {
	if nickname == "" {
		return nil, ErrPwdEmpty
	}

	a.Nickname = nickname
	return a, nil
}

// SetNickname 设置昵称
func (a *Account) SetAvatar(avatar string) (*Account, *Error) {
	if avatar == "" {
		return nil, ErrAvatarEmpty
	}
	a.Avatar = avatar

	return a, nil
}

// 设置角色
func (a *Account) SetRoles(roles []uuid.UUID) (*Account, *Error) {

	if len(roles) == 0 {
		return nil, ErrUserRoleFailed
	}

	return a, nil
}

// SetPassword 设置密码
func (a *Account) SetPassword(pwd string) (*Account, *Error) {
	if pwd == "" {
		return nil, ErrPwdEmpty
	}

	a.Pwd = pwd
	return a, nil
}

// DisableAccount 禁用账户
func (a *Account) DisableAccount() {
	a.Status = status.Disable
}

// EnableAccount 启用账户
func (a *Account) EnableAccount() {
	a.Status = status.Active
}

// IsEnabled 检查账户是否启用
func (a *Account) IsEnabled() bool {
	return a.Status == status.Active
}

// CheckPassword 检查密码是否正确
func (a *Account) CheckPassword(pwd string) bool {
	// TODO: 密码加密对比
	return a.Pwd == pwd
}
