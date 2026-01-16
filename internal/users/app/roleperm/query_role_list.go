package roleperm

import (
	"context"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	"gorm.io/gorm"
)

type RoleDto struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
	Name string    `json:"name"`
}

type RoleListQuery struct {
}

type RoleListQueryHandler struct {
	*gorm.DB
}

func NewRoleListQueryHandler(repo *gorm.DB) *RoleListQueryHandler {
	return &RoleListQueryHandler{
		DB: repo,
	}
}

func (h *RoleListQueryHandler) Handle(ctx context.Context, query RoleListQuery) ([]RoleDto, error) {

	var roles []roleperm.Role

	if tx := h.Find(&roles); tx.Error != nil {
		return nil, tx.Error
	}

	var dtos []RoleDto

	for _, role := range roles {
		dtos = append(dtos, RoleDto{
			ID:   role.ID,
			Role: role.Role,
			Name: role.Name,
		})
	}

	return dtos, nil
}
