package libError

import (
	"app/infrastructure/http/common"

	"github.com/go-errors/errors"
)

func IsAcceptable(target error, exceptions ...error) bool {

	if target == nil {

		return true
	}

	for _, exceptErr := range exceptions {

		if exceptErr == nil {

			continue
		}

		if target == exceptErr {

			return true
		}

	}

	return false
}

func NewInternal(errList ...error) error {

	return errors.Join(append(errList, common.ERR_INTERNAL)...)
}
