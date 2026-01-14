package roleperm

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/status"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

// FuncType 功能类型
type Type int8

const (
	Service  Type = iota + 1 // 服务
	Model                    // 模块
	Menu                     // 菜单
	MenuItem                 // 菜单项
	Action                   // 操作
)

// Permission 功能
type RolePermission struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Name                         string        // 名称
	Type                         Type          // 功能类型
	Route                        string        // 路由
	Icon                         string        // 图标
	Desc                         string        // 描述
	Weight                       int32         // 权重
	ParentID                     *uuid.UUID    // 父级ID
	Status                       status.Status // 状态
}

// newPermission 创建功能
func newPermission(name, route, icon, desc string, weight int32, ptype Type) (*RolePermission, *Error) {

	fun := &RolePermission{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Type:          ptype,
		Weight:        weight,
	}

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

	return fun, nil

}

// SetName 设置名称
func (f *RolePermission) SetName(name string) *Error {
	if name == "" {
		return ErrNameEmpty
	}
	f.Name = name
	return nil
}

// SetRoute 设置路由
func (f *RolePermission) SetRoute(route string) *Error {
	if route == "" {
		return ErrRouteEmpty
	}
	f.Route = route
	return nil
}

// SetIcon 设置图标
func (f *RolePermission) SetIcon(icon string) *Error {
	if icon == "" {
		return ErrIconEmpty
	}
	f.Icon = icon
	return nil
}

// SetDesc 设置描述
func (f *RolePermission) SetDesc(desc string) *Error {
	if desc == "" {
		return ErrDescEmpty
	}
	f.Desc = desc
	return nil
}
