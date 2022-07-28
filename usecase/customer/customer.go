package customer

import (
	"jamtangan/domain"
)

type customerUseCase struct {
	productRepository     domain.ProductRepository
	transactionRepository domain.TransactionRepository
}

func NewUseCase(
	productRepository domain.ProductRepository,
	transactionRepository domain.TransactionRepository,
) *customerUseCase {
	return &customerUseCase{
		productRepository:     productRepository,
		transactionRepository: transactionRepository,
	}
}
