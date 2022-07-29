package domain

import (
	"github.com/bwmarrin/snowflake"
)

type Snowflake interface {
	Generate() snowflake.ID
}
