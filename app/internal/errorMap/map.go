package errorMap

import (
	"app/internal/common"
	"errors"
	"fmt"
)

type ERROR_CODE_T = uint32

const (
	internal_bit_shift int = -iota + 31
	http_bit_shift
)

const (
	err_internal_flag ERROR_CODE_T = iota + 0x1<<internal_bit_shift
)

const (
	err_auth_flag ERROR_CODE_T = iota + 0x1<<http_bit_shift
	ERR_AUTH_FORCE_ANNONYMOUS
	ERR_AUTH_NO_ACCESSTOKEN
)

var (
	error_map map[ERROR_CODE_T]error = map[ERROR_CODE_T]error{
		ERR_AUTH_FORCE_ANNONYMOUS: errors.Join(
			common.ERR_FORBIDEN, code_error_t{ERR_AUTH_FORCE_ANNONYMOUS}, fmt.Errorf("you must logged out to access this route"),
		),
		ERR_AUTH_NO_ACCESSTOKEN: errors.Join(
			common.ERR_UNAUTHORIZED, code_error_t{ERR_AUTH_NO_ACCESSTOKEN}, fmt.Errorf("(%X) unauthorized", ERR_AUTH_NO_ACCESSTOKEN),
		),
	}
)

func Get(errCode ERROR_CODE_T) error {

	return error_map[errCode]
}
