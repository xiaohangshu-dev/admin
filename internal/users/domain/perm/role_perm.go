package perm

import "github.com/google/uuid"

// RolePerm 角色功能关系
type RolePerm struct {
	RoleId uuid.UUID
	PermId uuid.UUID
}

func newRolePerm(roleId, funcId uuid.UUID) RolePerm {
	return RolePerm{
		RoleId: roleId,
		PermId: funcId,
	}
}
