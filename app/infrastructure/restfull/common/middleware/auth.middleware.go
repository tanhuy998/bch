package middleware

import (
	"app/infrastructure/restfull/common/middleware/hook"
	"errors"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

const (
	AUTH_USER                = "auth_user"
	AUTH_HEADER              = "AUTH_HEADER"
	JWT_PUBLIC_KEY           = "JWT_PUBLIC_KEY"
	AUTH_REQ_HEADER          = "Authorization"
	AUTH_COOKIE_ACCESS_TOKEN = "access-token"
	AUTH_PASSED              = "auth-passed"
)

const (
	CTX_ACCESS_TOKEN_KEY = "ctx_access_token"
)

var (
	ERR_INVALID_ACCESS_TOKEN = errors.New("invalid access token")
	ERR_MISSING_AUTH_HEADER  = errors.New("missing authorization header")
	ERR_ACCESS_TOKEN_EXPIRED = errors.New("access token expired")
)

type (
	IErrorHandler interface {
		HandleContextError(iris.Context, error)
	}
)

func Auth(
	constraints ...hook.AuthorityConstraint,
) ContainerDependentMiddleware {

	return func(container *hero.Container) context.Handler {

		return container.Handler(
			func(
				ctx iris.Context,
				handler *AuthenticateMiddlewareHandler,
			) {

				handler.Execute(ctx, constraints)
			},
		)
	}
}
