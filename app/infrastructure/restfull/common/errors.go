package common

import "errors"

var (
	ERR_INVALID_HTTP_INPUT = errors.New("invalid input")
	ERR_INTERNAL           = errors.New("internal error")
	ERR_BAD_REQUEST        = errors.New("bad request")
	ERR_HTTP_NOT_FOUND     = errors.New("not found")
)
