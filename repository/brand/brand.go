package brand

import (
	"database/sql"

	"jamtangan/domain"
)

type brandRepository struct {
	sqlDB     *sql.DB
	snowflake domain.Snowflake
}

func NewRepository(
	sqlDB *sql.DB,
	snowflake domain.Snowflake,
) *brandRepository {
	return &brandRepository{
		sqlDB:     sqlDB,
		snowflake: snowflake,
	}
}
