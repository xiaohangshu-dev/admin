package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

type Status int

const (
	Normal  Status = iota // 正常
	Disable               // 禁用
)

// Account 描述用户账户领域对象
type Account struct {
	ddd.AggregateRoot[uuid.UUID]
	Phone       string       // 手机号
	Pwd         string       // 密码
	Salt        string       // 密码盐值
	Status      Status       // 状态
	CreatedAt   time.Time    // 创建时间
	Createby    *string      // 创建人(可空)
	UpdatedAt   *time.Time   // 更新时间(可空)
	UpdatedBy   *string      // 更新人(可空)
	UserProfile *UserProfile // 用户信息
}

// newAccount 创建账户并返回实例.
func newAccount(id uuid.UUID, phone, pwd string) (*Account, error) {

	if phone == "" {
		return nil, ErrPhoneEmpty
	}

	if pwd == "" {
		return nil, ErrPasswordEmpty
	}

	return &Account{
		AggregateRoot: ddd.NewAggregateRoot(id),
		Phone:         phone,
		Pwd:           pwd,
		Status:        Normal,
		CreatedAt:     time.Now(),
	}, nil
}

// DisableAccount 禁用账户
func (a *Account) DisableAccount() {
	a.Status = Disable
}

// EnableAccount 启用账户
func (a *Account) EnableAccount() {
	a.Status = Normal
}

// IsEnabled 检查账户是否启用
func (a *Account) IsEnabled() bool {
	return a.Status == Normal
}

// SetPassword 设置密码
func (a *Account) SetPassword(pwd string) {
	a.Pwd = pwd
}

// CheckPassword 检查密码是否正确
func (a *Account) CheckPassword(pwd string) bool {
	return a.Pwd == pwd
}

func (a *Account) UpdateUserProfile(profile *UserProfile) {

	a.UserProfile = profile
}
