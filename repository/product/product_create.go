package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"

	"jamtangan/domain"
)

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	const query = `INSERT INTO product 
		SET id = ?,
			name = ?,
			price = ?,
			brand_id = ?
		`

	if product.ID == 0 {
		product.ID = r.snowflake.Generate().Int64()
	}

	stmt, err := r.sqlDB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, product.ID, product.Name, product.Price, product.BrandID)
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) && mySQLError.Number == 1452 {
			return fmt.Errorf("brand ID not found: %w", domain.ErrBadRequest)
		}

		return err
	}

	return nil
}
