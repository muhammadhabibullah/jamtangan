package transaction

import (
	"context"

	"jamtangan/domain"
)

func (r *transactionRepository) GetByID(ctx context.Context, id int64) (domain.Transaction, []domain.TransactionProduct, error) {
	const (
		transactionQuery = `SELECT id, total_price, created_at, updated_at, is_deleted, deleted_at
			FROM transaction 
			WHERE id = ?`

		transactionProductQuery = `SELECT transaction_id, product_id, quantity, price, created_at, updated_at, is_deleted, deleted_at
			FROM transaction_product 
			WHERE transaction_id = ?`
	)

	rows, err := r.sqlDB.QueryContext(ctx, transactionQuery, id)
	if err != nil {
		return domain.Transaction{}, nil, err
	}
	defer rows.Close()

	transactions := make([]domain.Transaction, 0)
	for rows.Next() {
		var transaction domain.Transaction
		err = rows.Scan(
			&transaction.ID,
			&transaction.TotalPrice,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.IsDeleted,
			&transaction.DeletedAt,
		)
		if err != nil {
			return domain.Transaction{}, nil, err
		}

		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return domain.Transaction{}, nil, domain.ErrNotFound
	}
	var transaction = transactions[0]

	rows, err = r.sqlDB.QueryContext(ctx, transactionProductQuery, transaction.ID)
	if err != nil {
		return domain.Transaction{}, nil, err
	}
	defer rows.Close()

	transactionProducts := make([]domain.TransactionProduct, 0)
	for rows.Next() {
		var transactionProduct domain.TransactionProduct
		err = rows.Scan(
			&transactionProduct.TransactionID,
			&transactionProduct.ProductID,
			&transactionProduct.Quantity,
			&transactionProduct.Price,
			&transactionProduct.CreatedAt,
			&transactionProduct.UpdatedAt,
			&transactionProduct.IsDeleted,
			&transactionProduct.DeletedAt,
		)
		if err != nil {
			return domain.Transaction{}, nil, err
		}

		transactionProducts = append(transactionProducts, transactionProduct)
	}

	return transaction, transactionProducts, err
}
