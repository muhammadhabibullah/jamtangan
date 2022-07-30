package domain

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateTransactionRequest struct {
	TransactionProducts []CreateTransactionProductRequest `json:"transaction_products"`
}

func (req CreateTransactionRequest) Validate() error {
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.TransactionProducts, validation.Required),
	); err != nil {
		return err
	}

	for _, tp := range req.TransactionProducts {
		if err := tp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (req CreateTransactionRequest) ToTransactionDetail() TransactionDetail {
	transactionProducts := make([]*TransactionProduct, len(req.TransactionProducts), len(req.TransactionProducts))
	for i, tp := range req.TransactionProducts {
		transactionProducts[i] = &TransactionProduct{
			ProductID: tp.ProductID,
			Quantity:  tp.Quantity,
		}
	}

	return TransactionDetail{
		TransactionProducts: transactionProducts,
	}
}

type CreateTransactionProductRequest struct {
	ProductID int64 `json:"product_id,string" example:"1552703849368653824"`
	Quantity  int64 `json:"quantity" example:"1"`
}

func (req CreateTransactionProductRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.ProductID, validation.Required),
		validation.Field(&req.Quantity, validation.Required),
	)
}

type TransactionDetail struct {
	Transaction         *Transaction          `json:"transaction,omitempty"`
	TransactionProducts []*TransactionProduct `json:"transaction_products"`
}

type CustomerUseCase interface {
	CreateTransaction(ctx context.Context, request *TransactionDetail) error
	GetTransactionByID(ctx context.Context, id int64) (TransactionDetail, error)
	FetchProductByBrandID(ctx context.Context, brandID int64) ([]Product, error)
}
