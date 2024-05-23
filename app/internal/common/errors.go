package common

import "errors"

var (
	ERR_INVALID_HTTP_INPUT = errors.New("invalid input")
	ERR_INTERNAL           = errors.New("internal error")
)
