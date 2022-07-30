package domain

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateBrandRequest struct {
	Name string `json:"name" example:"CASIO"`
}

func (req CreateBrandRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
	)
}

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
