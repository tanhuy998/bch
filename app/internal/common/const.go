package common

import (
	"fmt"
)

var (
	ERR_NOT_FOUND    = fmt.Errorf("(not found)")
	ERR_BAD_REQUEST  = fmt.Errorf("(bad request)")
	ERR_UNAUTHORIZED = fmt.Errorf("(unauthorized)")
	ERR_FORBIDEN     = fmt.Errorf("(forbiden)")
	ERR_INTERNAL     = fmt.Errorf("(internal error)")
	ERR_CONFLICT     = fmt.Errorf("(conflit)")
	ERR_TIMEOUT      = fmt.Errorf("(timeout)")
)
