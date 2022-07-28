package brand

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"

	"jamtangan/domain"
)

func (r *brandRepository) Create(ctx context.Context, brand *domain.Brand) error {
	const query = `INSERT INTO brand 
		SET id = ?,
			name = ?
		`

	if brand.ID == 0 {
		brand.ID = r.snowflake.Generate().Int64()
	}

	stmt, err := r.sqlDB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, brand.ID, brand.Name)
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) && mySQLError.Number == 1062 {
			return fmt.Errorf("%s: %w", err, domain.ErrDuplicate)
		}

		return err
	}

	return nil
}
