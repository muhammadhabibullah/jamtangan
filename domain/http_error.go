package domain

import (
	"encoding/json"
)

type HTTPError struct {
	Err string `json:"error"`
}

func NewHTTPError(err error) string {
	httpErr := HTTPError{
		Err: err.Error(),
	}

	httpErrorStr, _ := json.Marshal(&httpErr)
	return string(httpErrorStr)
}
