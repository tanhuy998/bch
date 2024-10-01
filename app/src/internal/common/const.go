package common

import "errors"

var (
	ERR_NOT_FOUND    = errors.New("not found:")
	ERR_BAD_REQUEST  = errors.New("bad request:")
	ERR_UNAUTHORIZED = errors.New("unauthorized:")
	ERR_FORBIDEN     = errors.New("forbiden:")
	ERR_INTERNAL     = errors.New("internal error")
)
