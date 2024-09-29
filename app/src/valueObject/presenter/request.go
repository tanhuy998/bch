package requestPresenter

import (
	"app/internal/common"
	"fmt"
	"regexp"
	"time"

	"github.com/kataras/iris/v12"
)

const (
	CANDIDATE_SIGNING_OLD = 17
	PARENT_THRESHOLD      = 18
	alert_name_msg        = `invalid %s, contains special characters, received "%s"`
)

var (
	// match unicode letters and whiteSpace that belong to formal names
	regex_match_name *regexp.Regexp = regexp.MustCompile(`^[\p{L}\s]{3,}$`)
)

type IRequestBinder interface {
	/*
		Structs that implement IRequestBinder define its own
		request context to add extra business logic after data
		transfered from request context to the request presenter object.
	*/
	Bind(ctx iris.Context) error
}

func isValidName(name string) bool {

	return regex_match_name.MatchString(name)
}

func errorAlert(key string, value string) error {

	return fmt.Errorf(alert_name_msg, key, value)
}

func validateCandidateDateOfBirth(date time.Time) error {

	currentYear := time.Now().Year()

	if currentYear-date.Year() < CANDIDATE_SIGNING_OLD {

		return common.ERR_INVALID_HTTP_INPUT
	}

	return nil
}

func validateFormalName(name string) error {

	if !regex_match_name.Match([]byte(name)) {

		return common.ERR_INVALID_HTTP_INPUT
	}

	return nil
}
