package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Manager 负责用户账号领域对象的管理操作,封装数据库访问。
type Manager struct {
	*gorm.DB
}

// NewManager 返回用户领域服务实例
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		DB: db,
	}
}

// Create 创建一个新的用户账号,返回用户账号对象或错误信息。
func (m *Manager) Create(username, nickname, avatar, pwd string, phone, email *string, createBy uuid.UUID, gender Gender) (*Account, error) {

	// 内部业务规则校验
	account, err := newAccount(username, nickname, avatar, pwd, createBy, gender)

	if err != nil {
		return nil, err
	}

	// 无需校验的参数进行赋值
	account.Phone = phone
	account.Email = email

	// 外部业务规则校验
	if err := m.Where("username = ?", username).First(&Account{}).Error; err == nil {
		return nil, ErrUserAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if phone != nil {
		if err := m.Where("phone = ?", *phone).First(&Account{}).Error; err == nil {
			return nil, ErrPhoneAlreadyExist
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return account, nil
}

// Update 修改用户账号信息,返回修改后的用户账号或错误信息
func (m *Manager) Update(id uuid.UUID, nickname, avatar string, phone, email *string, updateBy uuid.UUID, gender Gender) (*Account, error) {

	account := &Account{}

	if err := m.Where("id = ?", id).First(account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 内部业务规则校验
	if err := account.SetNickname(nickname); err != nil {
		return account, err
	}

	if err := account.SetAvatar(avatar); err != nil {
		return account, err
	}

	// 无需校验参数赋值
	account.Phone = phone
	account.Email = email
	account.Gender = gender
	account.UpdateBy = &updateBy

	// 外部业务规则校验
	if phone != nil {
		if err := m.Where("phone = ? AND id <> ?", *phone, id).First(&Account{}).Error; err == nil {
			return nil, ErrPhoneAlreadyExist
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return account, nil
}

// UpdatePwd 修改密码,返回修改密码后的账户信息或者错误
func (m *Manager) UpdatePwd(id uuid.UUID, pwd string, updateBy uuid.UUID) (*Account, error) {

	account := &Account{}

	if err := m.First(account, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if err := account.SetPassword(pwd); err != nil {
		return account, err
	}

	now := time.Now()
	account.UpdatedAt = &now
	account.UpdateBy = &updateBy

	return account, nil
}
