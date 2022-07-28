package admin

import (
	"context"
	"strings"

	"jamtangan/domain"
)

func (u *adminUseCase) CreateBrand(ctx context.Context, brandName string) (domain.Brand, error) {
	brand := domain.Brand{
		Name: strings.ToUpper(brandName),
	}

	err := u.brandRepository.Create(ctx, &brand)
	if err != nil {
		return domain.Brand{}, err
	}

	return u.brandRepository.GetByID(ctx, brand.ID)
}
