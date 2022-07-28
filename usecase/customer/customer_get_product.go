package customer

import (
	"context"

	"jamtangan/domain"
)

func (u *customerUseCase) FetchProductByBrandID(ctx context.Context, brandID int64) ([]domain.Product, error) {
	return u.productRepository.FetchByBrandID(ctx, brandID)
}
