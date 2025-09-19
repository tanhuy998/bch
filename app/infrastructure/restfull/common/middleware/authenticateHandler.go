package middleware

import (
	libAuth "app/infrastructure/restfull/common/auth/lib"
	"app/infrastructure/restfull/common/middleware/hook"
	"app/internal/errorMap"
	"app/port/authenticatorServicePort"

	"github.com/kataras/iris/v12"
)

type (
	AuthenticateMiddlewareHandler struct {
		Authenticator authenticatorServicePort.IAuthenticator
		ErrorHandler  IErrorHandler
	}
)

func (this *AuthenticateMiddlewareHandler) Execute(ctx iris.Context, authConstraints []hook.AuthorityConstraint) {

	accessToken, err := this.Authenticator.VerifyAccessToken(ctx)

	if err != nil {

		this.ErrorHandler.HandleContextError(ctx, err)
		return
	}

	switch {
	case accessToken != nil && isAuthExcludedPath(ctx):
		this.ErrorHandler.HandleContextError(
			// ctx, errors.Join(
			// 	common.ERR_FORBIDEN, fmt.Errorf("you must logged out to access this route"),
			// ),
			ctx, errorMap.Get(errorMap.ERR_AUTH_FORCE_ANNONYMOUS),
		)
		return
	case accessToken == nil && !isAllowAnonymousPath(ctx):
		this.ErrorHandler.HandleContextError(
			// ctx, errors.Join(
			// 	common.ERR_UNAUTHORIZED, fmt.Errorf("unauthorized"),
			// ),
			ctx, errorMap.Get(errorMap.ERR_AUTH_NO_ACCESSTOKEN),
		)
		return
	}

	switch {
	case accessToken == nil || len(authConstraints) == 0:
		ctx.Next()
		return
	}

	ctx.Values().Set(CTX_ACCESS_TOKEN_KEY, accessToken)

	err = this.Authenticator.Authorize(ctx)

	if err != nil {

		this.ErrorHandler.HandleContextError(ctx, err)
		return
	}

	ctx.Next()
}

/*
Check whether the the request path is excluded from auth

example: path /api needs authentication but it's child path /api/login doesn't need auth
*/
func isAuthExcludedPath(ctx iris.Context) bool {

	return libAuth.CheckAuthExCludedPath(
		ctx.Method() + ctx.Path(),
	)
}

func isAllowAnonymousPath(ctx iris.Context) bool {

	return libAuth.CheckAuthAllowAnonymouse(
		ctx.Method() + ctx.Path(),
	)
}
