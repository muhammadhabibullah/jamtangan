package domain

import (
	"fmt"
)

var ErrNotFound = fmt.Errorf("not found")

var ErrDuplicate = fmt.Errorf("duplicate")

var ErrBadRequest = fmt.Errorf("bad request")

var ErrInvalidID = fmt.Errorf("invalid ID")

var ErrInvalidRequestMethod = fmt.Errorf("invalid request method")
