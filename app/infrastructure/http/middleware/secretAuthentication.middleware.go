package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/kataras/iris/v12"
)

const (
	SECRET_AUTH_SCHEME = "Secret"
	SECRET_AUTH_REALMM = "tenant"
	ENV_APP_SECRET     = "APP_SECRET"
)

type (
	default_reponse struct {
		Message string `json:"message"`
	}
)

var (
	regexp_sep_realm    = regexp.MustCompile(`realm="tenant"`)
	regexp_sep_response = regexp.MustCompile(
		regexp.QuoteMeta(
			fmt.Sprintf(`response="%s"`, os.Getenv(ENV_APP_SECRET)),
		),
	)
)

func SecretAuth(ctx iris.Context) {

	auth_header_val := ctx.GetHeader("Authorization")

	if auth_header_val == "" {

		responseChallenges(ctx)
		return
	}

	if validateSecretAuth(auth_header_val) != nil {

		responseChallenges(ctx)
		return
	}

	ctx.Next()
}

func validateSecretAuth(value string) error {

	defaultErr := errors.New("401 unauthorizaed")

	if !strings.HasPrefix(value, fmt.Sprintf(`%s `, SECRET_AUTH_SCHEME)) {

		return defaultErr
	}

	temp := strings.Split(value, " ")

	if len(temp) < 2 {

		return defaultErr
	}

	p := temp[1]

	params := strings.Split(p, ",")

	if len(params) < 2 {

		return defaultErr
	}

	if !regexp_sep_realm.MatchString(temp[1]) {

		return defaultErr
	}

	if !regexp_sep_response.MatchString(temp[1]) {

		return defaultErr
	}

	return nil
}

func responseChallenges(ctx iris.Context) {

	defer ctx.EndRequest()

	//common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusUnauthorized, "401 unauthorized")
	ctx.StatusCode(http.StatusUnauthorized)
	ctx.ResponseWriter().
		Header().
		Add(
			"WWW-Authenticate", fmt.Sprintf(`%s realm="%s"`, SECRET_AUTH_SCHEME, SECRET_AUTH_REALMM),
		)

	_ = ctx.JSON(default_reponse{Message: "401 unauthorized"})

	ctx.EndRequest()
}
