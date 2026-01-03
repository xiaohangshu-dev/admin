package user

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/gender"
	"gorm.io/gorm"
)

// Manager 负责用户账号领域对象的管理操作,封装数据库访问。
type Manager struct {
	*gorm.DB
}

// NewUserManager 返回用户领域服务实例
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		DB: db,
	}
}

// Create 创建一个新的用户账号,返回用户账号对象或错误信息。
func (m *Manager) Create(username, nickname, avatar, pwd string, phone, email *string, createBy uuid.UUID, gender gender.Gender, roles []uuid.UUID) (*Account, *Error) {

	// 外部业务规则校验
	account := &Account{}
	m.Where("username = ?", username).First(account)

	// 用户名唯一校验
	if account.ID != uuid.Nil {
		return nil, ErrUserAlreadyExists
	}

	// 手机号唯一校验
	if phone != nil {
		m.Where("phone = ?", phone).First(account)

		if account.ID != uuid.Nil {
			return nil, ErrPhoneAlreadyExist
		}
	}

	// 内部业务规则校验
	account, err := newAccount(uuid.New(), username, nickname, avatar, pwd, createBy, gender, roles)

	if err != nil {
		return nil, err
	}

	// 无需校验的参数进行赋值
	account.Phone = phone
	account.Email = email

	return account, nil
}

// Update 修改用户账号信息,返回修改后的用户账号或错误信息
func (m *Manager) Update(id uuid.UUID, nickname, avatar string, phone, email *string, updateBy uuid.UUID, gender gender.Gender, roles []uuid.UUID) (*Account, *Error) {

	account := &Account{}
	m.First(account, id)

	// 用户是否存在
	if account.ID != uuid.Nil {
		return nil, ErrUserNotFound
	}

	// 手机号唯一校验
	if phone != nil {
		m.Where("phone = ? && id != ?", phone, id).First(account)

		if account.ID != uuid.Nil {
			return nil, ErrPhoneAlreadyExist
		}
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

	// 无需校验的参数进行赋值
	account.Phone = phone
	account.Email = email

	return account, nil
}
