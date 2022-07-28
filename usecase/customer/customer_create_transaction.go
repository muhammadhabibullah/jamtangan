package customer

import (
	"context"
	"fmt"

	"jamtangan/domain"
)

func (u *customerUseCase) CreateTransaction(
	ctx context.Context,
	request *domain.TransactionDetail,
) error {
	if len(request.TransactionProducts) == 0 {
		return fmt.Errorf("no product to buy: %w", domain.ErrBadRequest)
	}

	var totalPrice int64
	for i, transactionProduct := range request.TransactionProducts {
		product, err := u.productRepository.GetByID(ctx, transactionProduct.ProductID)
		if err != nil {
			return err
		}

		request.TransactionProducts[i].Price = product.Price
		totalPrice += transactionProduct.Quantity * product.Price
	}

	request.Transaction = &domain.Transaction{
		TotalPrice: totalPrice,
	}
	return u.transactionRepository.Create(ctx, request.Transaction, request.TransactionProducts)
}
