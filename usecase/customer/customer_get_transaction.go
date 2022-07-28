package customer

import (
	"context"

	"jamtangan/domain"
)

func (u *customerUseCase) GetTransactionByID(ctx context.Context, id int64) (domain.TransactionDetail, error) {
	transaction, transactionProducts, err := u.transactionRepository.GetByID(ctx, id)
	if err != nil {
		return domain.TransactionDetail{}, err
	}
	if len(transactionProducts) == 0 {
		return domain.TransactionDetail{}, nil
	}

	transactionDetail := domain.TransactionDetail{
		Transaction:         &transaction,
		TransactionProducts: make([]*domain.TransactionProduct, len(transactionProducts), len(transactionProducts)),
	}
	for i := range transactionProducts {
		transactionDetail.TransactionProducts[i] = &transactionProducts[i]
	}

	return transactionDetail, nil
}
