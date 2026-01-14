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
type Permission struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	ParentID                     *uuid.UUID    // 父级ID
	Type                         Type          // 功能类型
	Title                        string        // 名称
	Perm                         string        // 权限标识
	Route                        string        // 路由
	Icon                         string        // 图标
	Desc                         string        // 描述
	Weight                       int32         // 权重
	Status                       status.Status // 状态
}

// newPermission 创建功能
func newPermission(name, route, icon, desc string, weight int32, ptype Type) (*Permission, *Error) {

	fun := &Permission{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Type:          ptype,
		Weight:        weight,
	}

	if err := fun.SetTitle(name); err != nil {
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

// SetTitle 设置标题
func (f *Permission) SetTitle(title string) *Error {
	if title == "" {
		return ErrNameEmpty
	}
	f.Title = title
	return nil
}

// SetPerm 设置权限标识
func (f *Permission) SetPerm(perm string) *Error {
	if perm == "" {
		return ErrPermEmpty
	}
	f.Perm = perm
	return nil
}

// SetRoute 设置路由
func (f *Permission) SetRoute(route string) *Error {
	if route == "" {
		return ErrRouteEmpty
	}
	f.Route = route
	return nil
}

// SetIcon 设置图标
func (f *Permission) SetIcon(icon string) *Error {
	if icon == "" {
		return ErrIconEmpty
	}
	f.Icon = icon
	return nil
}

// SetDesc 设置描述
func (f *Permission) SetDesc(desc string) *Error {
	if desc == "" {
		return ErrDescEmpty
	}
	f.Desc = desc
	return nil
}
