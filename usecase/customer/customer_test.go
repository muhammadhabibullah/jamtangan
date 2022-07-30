package customer_test

import (
	"testing"

	"github.com/onsi/gomega"

	"jamtangan/domain"
	mocks "jamtangan/mock"
	"jamtangan/usecase/customer"
)

type customerTest struct {
	customerUseCase       domain.CustomerUseCase
	productRepository     *mocks.ProductRepository
	transactionRepository *mocks.TransactionRepository
}

func test(t *testing.T, fn func(*gomega.WithT, *customerTest)) {
	productRepoMock := new(mocks.ProductRepository)
	transactionRepoMock := new(mocks.TransactionRepository)
	customerUseCase := customer.NewUseCase(productRepoMock, transactionRepoMock)

	a := customerTest{
		customerUseCase:       customerUseCase,
		productRepository:     productRepoMock,
		transactionRepository: transactionRepoMock,
	}

	g := gomega.NewGomegaWithT(t)
	fn(g, &a)
}
