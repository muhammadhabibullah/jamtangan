package admin

import (
	"context"
	"fmt"

	"jamtangan/domain"
)

func (u *adminUseCase) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	if product.Price <= 0 {
		return domain.Product{}, fmt.Errorf("price must be exist: %w", domain.ErrBadRequest)
	}

	err := u.productRepository.Create(ctx, &product)
	if err != nil {
		return domain.Product{}, err
	}

	return u.productRepository.GetByID(ctx, product.ID)
}
