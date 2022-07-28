package admin

import (
	"context"

	"jamtangan/domain"
)

func (u *adminUseCase) GetProductByID(ctx context.Context, id int64) (domain.Product, error) {
	return u.productRepository.GetByID(ctx, id)
}
