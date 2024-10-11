package libError

import (
	"app/internal/common"
	"fmt"

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

	stack := errors.New("")

	errList = append([]error{common.ERR_INTERNAL}, errList...)
	errList = append(errList, fmt.Errorf(stack.ErrorStack()))

	return errors.Join(errList...)
}
