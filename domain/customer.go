package domain

import (
	"context"
)

type TransactionDetail struct {
	Transaction         *Transaction          `json:"transaction,omitempty"`
	TransactionProducts []*TransactionProduct `json:"transaction_products"`
}

type CustomerUseCase interface {
	CreateTransaction(ctx context.Context, request *TransactionDetail) error
	GetTransactionByID(ctx context.Context, id int64) (TransactionDetail, error)
	FetchProductByBrandID(ctx context.Context, brandID int64) ([]Product, error)
}
