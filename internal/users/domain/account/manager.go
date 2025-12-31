package account

import (
	"github.com/google/uuid"
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
func (m *Manager) Create(phone, password, createby string) (*Account, error) {

	account := &Account{}
	m.Where("phone = ?", phone).First(account)

	if account.ID != uuid.Nil {
		return nil, ErrUserAlreadyExists
	}

	account, err := newAccount(uuid.New(), phone, password)

	if err != nil {
		return nil, err
	}

	return account, nil
}
