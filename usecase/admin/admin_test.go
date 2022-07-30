package admin_test

import (
	"testing"

	"github.com/onsi/gomega"

	"jamtangan/domain"
	mocks "jamtangan/mock"
	"jamtangan/usecase/admin"
)

type adminTest struct {
	adminUseCase      domain.AdminUseCase
	brandRepository   *mocks.BrandRepository
	productRepository *mocks.ProductRepository
}

func test(t *testing.T, fn func(*gomega.WithT, *adminTest)) {
	brandRepoMock := new(mocks.BrandRepository)
	productRepoMock := new(mocks.ProductRepository)
	adminUseCase := admin.NewUseCase(brandRepoMock, productRepoMock)

	a := adminTest{
		adminUseCase:      adminUseCase,
		brandRepository:   brandRepoMock,
		productRepository: productRepoMock,
	}

	g := gomega.NewGomegaWithT(t)
	fn(g, &a)
}
