package function

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Manager 负责功能领域对象的管理操作,封装数据库访问。
type Manager struct {
	*gorm.DB
}

// NewUserManager 返回功能领域服务实例
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		DB: db,
	}
}

// Create 创建一个新的用户账号,返回用户账号对象或错误信息。
func (m *Manager) Create(name, route, icon, desc string, weight int32, ftype Type, parentID *uuid.UUID, createBy uuid.UUID) (*Function, error) {

	// 内部业务规则校验
	fun, err := newFunction(name, route, icon, desc, weight, ftype)

	if err != nil {
		return nil, err
	}

	// 无需校验的参数进行赋值
	fun.ParentID = parentID

	// 外部业务规则校验
	if err := m.Where("name = ? And parent_id= ?", name, parentID).First(&Function{}).Error; err == nil {
		return nil, ErrFunctionAlreadyExists
	}

	return fun, nil
}

// Update 修改用户账号信息,返回修改后的用户账号或错误信息
func (m *Manager) Update(id uuid.UUID, name, route, icon, desc string, weight int32, ftype Type, parentID *uuid.UUID, updateBy uuid.UUID) (*Function, error) {

	fun := &Function{}

	if err := m.Where("id = ?", id).First(fun).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFunctionNotFound
		}
		return nil, err
	}

	// 内部业务规则校验
	if err := fun.SetName(name); err != nil {
		return nil, err
	}

	if err := fun.SetRoute(route); err != nil {
		return nil, err
	}

	if err := fun.SetIcon(icon); err != nil {
		return nil, err
	}

	if err := fun.SetDesc(desc); err != nil {
		return nil, err
	}
	// 无需校验参数赋值
	now := time.Now()
	fun.ParentID = parentID
	fun.UpdatedAt = &now
	// 外部业务规则校验
	if err := m.Where("name = ? And parent_id= ?", name, parentID).First(&Function{}).Error; err == nil {
		return nil, ErrFunctionAlreadyExists
	}

	return fun, nil
}
