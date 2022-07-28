package brand

import (
	"database/sql"

	"github.com/bwmarrin/snowflake"
)

type brandRepository struct {
	sqlDB     *sql.DB
	snowflake *snowflake.Node
}

func NewRepository(
	sqlDB *sql.DB,
	snowflakeNode *snowflake.Node,
) *brandRepository {
	return &brandRepository{
		sqlDB:     sqlDB,
		snowflake: snowflakeNode,
	}
}
