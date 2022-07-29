package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"

	"jamtangan/domain"
)

func (r *transactionRepository) Create(
	ctx context.Context,
	transaction *domain.Transaction,
	transactionProducts []*domain.TransactionProduct,
) error {
	const (
		transactionQuery = `INSERT INTO transaction 
		SET id = ?,
			total_price = ?
		`

		transactionProductQuery = `INSERT INTO transaction_product 
		SET transaction_id = ?,
			product_id = ?,
			quantity = ?,
			price = ?
		`
	)

	if transaction.ID == 0 {
		transaction.ID = r.snowflake.Generate().Int64()
		for i := range transactionProducts {
			transactionProducts[i].TransactionID = transaction.ID
		}
	}

	tx, err := r.sqlDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, transactionQuery)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = stmt.ExecContext(ctx, transaction.ID, transaction.TotalPrice)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	stmt, err = tx.PrepareContext(ctx, transactionProductQuery)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, transactionProduct := range transactionProducts {
		_, err = stmt.ExecContext(ctx,
			transactionProduct.TransactionID,
			transactionProduct.ProductID,
			transactionProduct.Quantity,
			transactionProduct.Price,
		)
		if err != nil {
			var mySQLError *mysql.MySQLError
			if errors.As(err, &mySQLError) && mySQLError.Number == 1062 {
				err = fmt.Errorf("duplicate product ID: %w", domain.ErrBadRequest)
			}

			_ = tx.Rollback()
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
