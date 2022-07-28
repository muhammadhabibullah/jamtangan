package domain

import (
	"context"
	"time"
)

type Transaction struct {
	ID         int64      `json:"id,string"`
	TotalPrice int64      `json:"total_price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	IsDeleted  bool       `json:"is_deleted"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

type TransactionProduct struct {
	TransactionID int64      `json:"transaction_id,string"`
	ProductID     int64      `json:"product_id,string"`
	Quantity      int64      `json:"quantity"`
	Price         int64      `json:"price"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	IsDeleted     bool       `json:"is_deleted"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

type TransactionRepository interface {
	Create(context.Context, *Transaction, []*TransactionProduct) error
	GetByID(context.Context, int64) (Transaction, []TransactionProduct, error)
}
