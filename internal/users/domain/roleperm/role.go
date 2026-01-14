package roleperm

import (
	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/dic/status"
	"github.com/xiaohangshuhub/go-workit/pkg/ddd"
)

type Role struct {
	ddd.AggregateRoot[uuid.UUID]               // ID
	Role                         string        // 角色
	Name                         string        // 名称
	ParentID                     uuid.UUID     // 父级ID
	CreateBy                     string        // 创建人
	UpdateBy                     *string       // 更新人
	Status                       status.Status // 状态
}

// newRole 创建角色
func newRole(role, name, createBy string, parentID uuid.UUID) (*Role, *Error) {
	r := &Role{
		AggregateRoot: ddd.NewAggregateRoot(uuid.New()),
		Status:        status.Enable,
	}

	if r.SetRole(role) != nil {
		return nil, ErrNameEmpty
	}
	if r.SetName(name) != nil {
		return nil, ErrNameEmpty
	}
	if r.SetParentID(parentID) != nil {
		return nil, ErrNameEmpty
	}
	if createBy == "" {
		return nil, ErrNameEmpty
	}
	return r, nil
}

// SetRole 设置角色标识
func (r *Role) SetRole(role string) *Error {
	if role == "" {
		return ErrNameEmpty
	}
	r.Name = role
	return nil
}

// SetName 设置名称
func (r *Role) SetName(name string) *Error {
	if name == "" {
		return ErrNameEmpty
	}
	r.Name = name
	return nil
}

// SetParentID 设置父级ID
func (r *Role) SetParentID(parentID uuid.UUID) *Error {
	if parentID == uuid.Nil {
		return ErrNameEmpty
	}
	r.ParentID = parentID
	return nil
}
