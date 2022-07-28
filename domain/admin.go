package domain

import (
	"context"
)

type AdminUseCase interface {
	CreateBrand(ctx context.Context, brandName string) (Brand, error)
	CreateProduct(ctx context.Context, product Product) (Product, error)
	GetProductByID(ctx context.Context, id int64) (Product, error)
}
