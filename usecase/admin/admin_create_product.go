package admin

import (
	"context"

	"jamtangan/domain"
)

func (u *adminUseCase) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	err := u.productRepository.Create(ctx, &product)
	if err != nil {
		return domain.Product{}, err
	}

	return u.productRepository.GetByID(ctx, product.ID)
}
