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
	var totalPrice int64
	for i, transactionProduct := range request.TransactionProducts {
		product, err := u.productRepository.GetByID(ctx, transactionProduct.ProductID)
		if err != nil {
			return fmt.Errorf("product ID %d: %w", transactionProduct.ProductID, err)
		}

		request.TransactionProducts[i].Price = product.Price
		totalPrice += transactionProduct.Quantity * product.Price
	}

	request.Transaction = &domain.Transaction{
		TotalPrice: totalPrice,
	}
	err := u.transactionRepository.Create(ctx, request.Transaction, request.TransactionProducts)
	if err != nil {
		return err
	}

	*request, err = u.GetTransactionByID(ctx, request.Transaction.ID)
	return err
}
