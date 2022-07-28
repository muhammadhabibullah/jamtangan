package admin

import (
	"jamtangan/domain"
)

type adminUseCase struct {
	brandRepository   domain.BrandRepository
	productRepository domain.ProductRepository
}

func NewUseCase(
	brandRepository domain.BrandRepository,
	productRepository domain.ProductRepository,
) *adminUseCase {
	return &adminUseCase{
		brandRepository:   brandRepository,
		productRepository: productRepository,
	}
}
