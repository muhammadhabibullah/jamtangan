package domain

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TransactionDetail struct {
	Transaction         *Transaction          `json:"transaction,omitempty"`
	TransactionProducts []*TransactionProduct `json:"transaction_products"`
}

func (td TransactionDetail) Validate() error {
	if err := validation.ValidateStruct(&td,
		validation.Field(&td.TransactionProducts, validation.Required),
	); err != nil {
		return err
	}

	for _, tp := range td.TransactionProducts {
		if err := tp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CustomerUseCase interface {
	CreateTransaction(ctx context.Context, request *TransactionDetail) error
	GetTransactionByID(ctx context.Context, id int64) (TransactionDetail, error)
	FetchProductByBrandID(ctx context.Context, brandID int64) ([]Product, error)
}
