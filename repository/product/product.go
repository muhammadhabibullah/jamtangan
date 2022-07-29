package product

import (
	"database/sql"

	"jamtangan/domain"
)

type productRepository struct {
	sqlDB     *sql.DB
	snowflake domain.Snowflake
}

func NewRepository(
	sqlDB *sql.DB,
	snowflake domain.Snowflake,
) *productRepository {
	return &productRepository{
		sqlDB:     sqlDB,
		snowflake: snowflake,
	}
}
