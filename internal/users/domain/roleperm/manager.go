package roleperm

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Manager 负责管理用户权限相关的业务逻辑
type Manager struct {
	*gorm.DB
}

// NewManager 返回一个新的 Manager 实例
func NewManager(db *gorm.DB) *Manager {
	return &Manager{
		DB: db,
	}
}

// CreatePremission 创建权限
func (m *Manager) CreatePremission(name, route, icon, desc, createBy string, weight int32, ptype Type, parentID *uuid.UUID) (*Permission, error) {

	// 内部业务规则校验
	fun, err := newPermission(name, route, icon, desc, weight, ptype)

	if err != nil {
		return nil, err
	}

	// 无需校验的参数进行赋值
	fun.ParentID = parentID

	// 外部业务规则校验
	if err := m.Where("name = ? And parent_id= ?", name, parentID).First(&Permission{}).Error; err == nil {
		return nil, ErrFunctionAlreadyExists
	}

	return fun, nil
}

// updatePermission 更新权限
func (m *Manager) UpdatePermission(id uuid.UUID, title, route, icon, desc, updateBy string, weight int32, ptype Type, parentID *uuid.UUID) (*Permission, error) {

	perm := &Permission{}

	if err := m.Where("id = ?", id).First(perm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFunctionNotFound
		}
		return nil, err
	}

	// 内部业务规则校验
	if err := perm.SetTitle(title); err != nil {
		return nil, err
	}

	if err := perm.SetRoute(route); err != nil {
		return nil, err
	}

	if err := perm.SetIcon(icon); err != nil {
		return nil, err
	}

	if err := perm.SetDesc(desc); err != nil {
		return nil, err
	}
	// 无需校验参数赋值
	now := time.Now()
	perm.ParentID = parentID
	perm.UpdatedAt = &now
	perm.Type = ptype
	// 外部业务规则校验
	if err := m.Where("title = ? And parent_id= ?", title, parentID).First(&Permission{}).Error; err == nil {
		return nil, ErrFunctionAlreadyExists
	}

	return perm, nil
}

// CreateRole 创建角色
func (m *Manager) CreateRole(role, name, createBy string, parentID uuid.UUID) (*Role, error) {
	// 内部业务规则校验
	r, err := newRole(role, name, createBy, parentID)
	if err != nil {
		return nil, err
	}

	// 外部业务规则校验
	if err := m.Where("role = ?", role).First(&Role{}).Error; err == nil {
		return nil, ErrRoleAlreadyExists
	}

	return r, nil
}

// UpdateRole 更新角色
func (m *Manager) UpdateRole(id uuid.UUID, role, name, updateBy string, parentID uuid.UUID) (*Role, error) {

	// 查询角色
	r := &Role{}
	if err := m.Where("id = ?", id).First(r).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}

	// 内部业务规则校验
	if err := r.SetRole(role); err != nil {
		return nil, err
	}

	if err := r.SetName(name); err != nil {
		return nil, err
	}

	if err := r.SetParentID(parentID); err != nil {
		return nil, err
	}

	// 无需校验参数赋值
	now := time.Now()
	r.UpdatedAt = &now

	// 外部业务规则校验
	if err := m.Where("role = ?", role).First(&Role{}).Error; err == nil {
		return nil, ErrRoleAlreadyExists
	}

	return r, nil
}

// ConfigureUserRoles 配置用户角色
func (m *Manager) ConfigureUserRoles(userId uuid.UUID, roleIdList []uuid.UUID) error {

	var user_role []UserRole

	for _, roleId := range roleIdList {
		user_role = append(user_role, newUserRole(userId, roleId))
	}

	if result := m.Delete(&UserRole{}, "user_id =?", userId); result.Error != nil {
		return result.Error
	}

	if result := m.DB.Create(&user_role); result.Error != nil {
		return result.Error
	}

	return nil
}

// ConfigureRolePerms 配置角色权限
func (m *Manager) ConfigureRolePerms(roleId uuid.UUID, funcIdList []uuid.UUID) error {
	var role_func []RolePerm

	for _, permId := range funcIdList {
		role_func = append(role_func, newRolePerm(roleId, permId))
	}

	if result := m.Delete(&RolePerm{}, "role_id =?", roleId); result.Error != nil {
		return result.Error
	}

	if result := m.DB.Create(&role_func); result.Error != nil {
		return result.Error
	}

	return nil
}
