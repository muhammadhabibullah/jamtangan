package product

import (
	"database/sql"

	"github.com/bwmarrin/snowflake"
)

type productRepository struct {
	sqlDB     *sql.DB
	snowflake *snowflake.Node
}

func NewRepository(
	sqlDB *sql.DB,
	snowflakeNode *snowflake.Node,
) *productRepository {
	return &productRepository{
		sqlDB:     sqlDB,
		snowflake: snowflakeNode,
	}
}
