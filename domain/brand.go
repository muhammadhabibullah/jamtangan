package domain

import (
	"context"
	"time"
)

type Brand struct {
	ID        int64      `json:"id,string"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type BrandRepository interface {
	Create(context.Context, *Brand) error
	GetByID(ctx context.Context, id int64) (Brand, error)
}
