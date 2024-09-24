package middleware

import (
	"errors"
	"fmt"
	"maps"
	"net/http"
	"os"
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
	secret_auth_challenges = map[string]func(string) bool{
		"realm": func(v string) bool {

			v = strings.TrimPrefix(v, `"`)
			v = strings.TrimSuffix(v, `"`)

			return v == SECRET_AUTH_REALMM
		},
		"response": func(v string) bool {

			v = strings.TrimPrefix(v, `"`)
			v = strings.TrimSuffix(v, `"`)

			return v == os.Getenv(ENV_APP_SECRET)
		},
	}
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

	challenges := maps.Clone(secret_auth_challenges)

	for _, s := range params {

		p := strings.Split(s, "=")

		if len(p) < 2 {

			return defaultErr
		}

		if validate, ok := challenges[p[0]]; ok {

			if !validate(p[1]) {

				return defaultErr
			}

			delete(challenges, p[0])
		}
	}

	if len(challenges) != 0 {

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
