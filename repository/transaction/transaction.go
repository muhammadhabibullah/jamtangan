package transaction

import (
	"database/sql"

	"github.com/bwmarrin/snowflake"
)

type transactionRepository struct {
	sqlDB     *sql.DB
	snowflake *snowflake.Node
}

func NewRepository(
	sqlDB *sql.DB,
	snowflakeNode *snowflake.Node,
) *transactionRepository {
	return &transactionRepository{
		sqlDB:     sqlDB,
		snowflake: snowflakeNode,
	}
}
