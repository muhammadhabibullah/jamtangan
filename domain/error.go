package domain

import (
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")

var ErrDuplicate = fmt.Errorf("duplicate")

var ErrBadRequest = fmt.Errorf("bad request")

var ErrInvalidRequestMethod = fmt.Errorf("invalid request method")
