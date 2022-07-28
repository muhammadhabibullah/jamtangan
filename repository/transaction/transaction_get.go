package transaction

import (
	"context"

	"jamtangan/domain"
)

func (r *transactionRepository) GetByID(ctx context.Context, id int64) (domain.Transaction, []domain.TransactionProduct, error) {
	const (
		transactionQuery = `SELECT t.id, t.total_price, t.created_at, t.updated_at, t.is_deleted, t.deleted_at,
       		tp.transaction_id, tp.product_id, tp.quantity, tp.price, tp.created_at, tp.updated_at, tp.is_deleted, tp.deleted_at
			FROM transaction t
			INNER JOIN transaction_product tp on t.id = tp.transaction_id
			WHERE id = ?`
	)

	rows, err := r.sqlDB.QueryContext(ctx, transactionQuery, id)
	if err != nil {
		return domain.Transaction{}, nil, err
	}
	defer rows.Close()

	var (
		transaction         domain.Transaction
		transactionProducts = make([]domain.TransactionProduct, 0)
	)

	for rows.Next() {
		var transactionProduct domain.TransactionProduct

		err = rows.Scan(
			&transaction.ID,
			&transaction.TotalPrice,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.IsDeleted,
			&transaction.DeletedAt,
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
