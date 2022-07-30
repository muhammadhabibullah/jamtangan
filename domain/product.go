package domain

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateProductRequest struct {
	Name    string `json:"name" example:"Casio G-Shock GX-56BB-1DR King Kong Solar Powered WR 200M Black Resin Band"`
	Price   int64  `json:"price" example:"1450000"`
	BrandID int64  `json:"brand_id,string" example:"1552655170888798208"`
}

func (req CreateProductRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Price, validation.Required),
		validation.Field(&req.BrandID, validation.Required),
	)
}

func (req CreateProductRequest) ToProduct() Product {
	return Product{
		Name:    req.Name,
		Price:   req.Price,
		BrandID: req.BrandID,
	}
}

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
