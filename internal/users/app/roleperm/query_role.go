package roleperm

import (
	"context"

	"github.com/google/uuid"
	"github.com/xiaohangshu-dev/admin/internal/users/domain/roleperm"
	"gorm.io/gorm"
)

type RoleInfoDto struct {
	ID    uuid.UUID
	Name  string
	Role  string
	Perms []uuid.UUID
}

type RoleQuery struct {
	ID uuid.UUID
}

type RoleQueryHandler struct {
	*gorm.DB
}

func NewRoleQueryHandler(repo *gorm.DB) *RoleQueryHandler {
	return &RoleQueryHandler{
		DB: repo,
	}
}

func (h *RoleQueryHandler) Handle(ctx context.Context, query RoleQuery) (RoleInfoDto, error) {

	var role roleperm.Role

	if tx := h.First(&role, query.ID); tx.Error != nil {
		return RoleInfoDto{}, tx.Error
	}

	var role_perms []roleperm.RolePerm

	if tx := h.Where("role_id = ?", role.ID).Find(&role_perms); tx.Error != nil {
		return RoleInfoDto{}, tx.Error
	}

	var perms []uuid.UUID

	for _, role_perm := range role_perms {
		perms = append(perms, role_perm.PermId)
	}

	return RoleInfoDto{
		ID:    role.ID,
		Role:  role.Role,
		Name:  role.Name,
		Perms: perms,
	}, nil

}
