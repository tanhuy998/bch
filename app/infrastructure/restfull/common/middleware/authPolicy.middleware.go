package middleware

import (
	"app/internal/bootstrap"
	accessLogServicePort "app/port/accessLog"
	"app/service/noExpireTokenProvider"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

var (
	escaped_app_secret string = strings.ReplaceAll(
		strings.ReplaceAll(os.Getenv(bootstrap.ENV_APP_SECRET), `+`, `\+`),
		`/`, `\/`,
	)
)

var authSignaturePolicy struct {
	noExpireTokenProvider.NoExpireTokenProvider
}

var (
	regexp_sep_auth_policy = regexp.MustCompile(
		fmt.Sprintf(
			`^Secret \s*realm="auth-policies"\s*,\s*app-secret="%s"\s*$`,
			escaped_app_secret,
		),
	)
)

func AuthPolicies(container *hero.Container) iris.Handler {

	return container.Handler(
		func(
			accessLogger accessLogServicePort.IAccessLogger,
			ctx iris.Context,
		) {

			defer func() {

				ctx.Next()
			}()

			auth_header := ctx.GetHeader("Authorization")

			log.Default().Println(auth_header)

			if !regexp_sep_auth_policy.MatchString(auth_header) {

				return
			}

			logField := struct {
				Op       string        `json:"operation"`
				Policies []interface{} `json:"policies_applied"`
			}{"grant-auth-policy", make([]interface{}, 0)}

			req_queries := ctx.Request().URL.Query()

			if req_queries.Has("at-no-exp") {

				logField.Policies = append(logField.Policies, "at-no-exp")
				ctx.Values().Set(authSignaturePolicy.CtxNoExpireKey(), true)
			}

			accessLogger.PushTraceLogs(
				ctx, logField,
			)
		},
	)
}
