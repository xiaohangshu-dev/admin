package roleperm

import (
	"context"

	"github.com/google/uuid"
	"github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	"gorm.io/gorm"
)

type PermDto struct {
	ID       uuid.UUID  `json:"id"`
	Title    string     `json:"name"`
	Perm     string     `json:"perm"`
	Route    string     `json:"route"`
	Icon     string     `json:"icon"`
	Desc     string     `json:"desc"`
	Weight   int32      `json:"weight"`
	ParentID *uuid.UUID `json:"parent_id"`
}

type PermListQuery struct {
}

type PermListQueryHandler struct {
	DB *gorm.DB
}

func NewPermListQueryHandler(repo *gorm.DB) *PermListQueryHandler {
	return &PermListQueryHandler{
		DB: repo,
	}
}

func (h *PermListQueryHandler) Handle(ctx context.Context, query PermListQuery) ([]*PermDto, error) {

	var perms []roleperm.Permission

	if tx := h.DB.Find(&perms); tx.Error != nil {
		return nil, tx.Error
	}

	var dtos []*PermDto

	for _, perm := range perms {
		dtos = append(dtos, &PermDto{
			ID:       perm.ID,
			Title:    perm.Title,
			Perm:     perm.Perm,
			Route:    perm.Route,
			Icon:     perm.Icon,
			Desc:     perm.Desc,
			Weight:   perm.Weight,
			ParentID: perm.ParentID,
		})
	}

	return dtos, nil

}
