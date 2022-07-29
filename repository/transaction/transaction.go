package transaction

import (
	"database/sql"

	"jamtangan/domain"
)

type transactionRepository struct {
	sqlDB     *sql.DB
	snowflake domain.Snowflake
}

func NewRepository(
	sqlDB *sql.DB,
	snowflakeNode domain.Snowflake,
) *transactionRepository {
	return &transactionRepository{
		sqlDB:     sqlDB,
		snowflake: snowflakeNode,
	}
}
