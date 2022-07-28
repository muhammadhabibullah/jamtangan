package domain

import (
	"context"
	"time"
)

type Product struct {
	ID        int64      `json:"id,string"`
	Name      string     `json:"name"`
	Price     int64      `json:"price"`
	BrandID   int64      `json:"brand_id,string"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ProductRepository interface {
	Create(context.Context, *Product) error
	GetByID(context.Context, int64) (Product, error)
	FetchByBrandID(context.Context, int64) ([]Product, error)
}
