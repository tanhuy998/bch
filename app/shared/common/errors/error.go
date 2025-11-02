package errors

import (
	"errors"
)

type (
	IComparableError interface {
		IsOneOf(errors ...error) bool
		MatchOneOf(errs ...error) error
	}

	compare_error_t struct {
		error
	}
)

func (this compare_error_t) IsOneOf(errors ...error) bool {

	if this.error == nil {
		return false
	}

	return this.MatchOneOf(errors...) != nil
}

func (this compare_error_t) MatchOneOf(errs ...error) error {

	if this.error == nil {
		return nil
	}

	for _, e := range errs {

		if errors.Is(this, e) {

			return e
		}
	}

	return nil
}

func Compare(err error) IComparableError {

	return compare_error_t{error: err}
}
