package middleware

import (
	"app/infrastructure/http/common"
	libAuth "app/infrastructure/http/common/auth/lib"
	"app/infrastructure/http/middleware/middlewareHelper"
	libIris "app/internal/lib/iris"
	accessTokenServicePort "app/port/accessToken"
	accessTokenClientPort "app/port/accessTokenClient"
	usecasePort "app/port/usecase"
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
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

var (
	ERR_INVALID_ACCESS_TOKEN = errors.New("invalid access token")
	ERR_MISSING_AUTH_HEADER  = errors.New("missing authorization header")
	ERR_ACCESS_TOKEN_EXPIRED = errors.New("access token expired")
)

type ()

func Auth(
	container *hero.Container,
	constraints ...middlewareHelper.AuthorityConstraint,
) iris.Handler {

	// authenticate := container.Handler(authentication_func)

	// return func(ctx iris.Context) {

	// 	accessToken := libIris.GetAccessToken(ctx)

	// 	if accessToken == nil {
	// 		// do dependencies injection
	// 		//container.Handler(authentication_func)(ctx)
	// 		authenticate(ctx)
	// 	}

	// 	if ctx.IsStopped() {
	// 		/**
	// 		authentication will stop the excution of the handler chain when
	// 		the request context doesn't match the route's requirements
	// 		*/
	// 		return
	// 	}

	// 	accessToken = libIris.GetAccessToken(ctx)

	// 	if accessToken == nil {
	// 		/*
	// 			No access token but but the reqeust context was not stopped
	// 			means the current request context's path is auth excluded
	// 		*/
	// 		ctx.Next()
	// 		return
	// 	}

	// 	if len(constraints) == 0 {

	// 		ctx.Next()
	// 		return
	// 	}

	// 	if !validateAuthority(accessToken, constraints) {

	// 		libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusForbidden, "forbiden authority")
	// 		return
	// 	}

	// 	ctx.Next()
	// }

	return container.Handler(
		func(
			ctx iris.Context,
			accessTokenClient accessTokenClientPort.IAccessTokenClient,
			checkAuthoritySessionUseCase usecasePort.IMiddlewareUseCase,
			errorHandler common.IMiddlewareErrorHandler,
		) {

			err := authenticate(ctx, accessTokenClient, checkAuthoritySessionUseCase)

			if err != nil {

				errorHandler.HandleContextError(ctx, err)
				return
			}

			accessToken := libIris.GetAccessToken(ctx)

			if accessToken == nil {

				ctx.Next()
				return
			}

			if len(constraints) == 0 {

				ctx.Next()
				return
			}

			if !validateAuthority(accessToken, constraints) {

				libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusForbidden, "forbiden authority")
				return
			}

			ctx.Next()
		},
	)
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

func isAnonymousPath(ctx iris.Context) bool {

	return libAuth.CheckAuthAnonymouse(
		ctx.Method() + ctx.Path(),
	)
}

func validateAuthority(accessToken accessTokenServicePort.IAccessToken, constraints []middlewareHelper.AuthorityConstraint) bool {

	if accessToken == nil {

		return false
	}

	for _, f := range constraints {

		if !f(accessToken) {

			return false
		}
	}

	return true
}

/*
This function parsed request's access token and checks whether
+ if there is no access token: check if the request's is anonymous route or auth exluded route.
+ if there is valid access token: check if the user session is known by the server.
+ otherwise: the request context will be abort and error code will be sent as response.

This func argument is injected by the dependency injection container
*/
func authentication_func(
	ctx iris.Context,
	accessTokenClient accessTokenClientPort.IAccessTokenClient,
	checkAuthoritySessionUseCase usecasePort.IMiddlewareUseCase,
	errorHandler common.IMiddlewareErrorHandler,
) {

	accessToken, err := accessTokenClient.Read(ctx)

	// switch {
	// case err == nil:
	// // case jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH:
	// // 	common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, err.Error())
	// case errors.Is(err, internalCommon.ERR_INTERNAL):
	// 	libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, "internal error")
	// default:
	// 	libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusUnauthorized, "(Authentication middleware error) unauthorized")
	// }

	// if err != nil {

	// 	return
	// }

	if err != nil {

		errorHandler.HandleContextError(ctx, err)
		return
	}

	isExcludedPath := isAuthExcludedPath(ctx)
	isAnonymousPath := isAnonymousPath(ctx)

	if accessToken != nil {

		ctx.Values().Set(common.CTX_ACCESS_TOKEN_KEY, accessToken)
	}

	if isExcludedPath && accessToken != nil {

		//libIris.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusBadRequest, "bad request")
		errorHandler.HandleContextError(ctx, errors.New("bad request"))
		return
	}

	if isAnonymousPath {

		return
	}

	if accessToken == nil {

		return
	}

	err = checkAuthoritySessionUseCase.Execute(ctx)

	if err != nil {

		ctx.Values().Remove(common.CTX_ACCESS_TOKEN_KEY)

		// if errors.Is(err, internalCommon.ERR_INTERNAL) {

		// 	libIris.SendDefaulJsonBodyAndEndRequest(ctx, 500, err.Error())
		// 	return
		// }

		// libIris.SendDefaulJsonBodyAndEndRequest(ctx, 401, err.Error())
		// return

		errorHandler.HandleContextError(ctx, err)
	}
}

func authenticate(
	ctx iris.Context,
	accessTokenClient accessTokenClientPort.IAccessTokenClient,
	checkAuthoritySessionUseCase usecasePort.IMiddlewareUseCase,
) error {

	accessToken, err := accessTokenClient.Read(ctx)

	if err != nil {

		return err
	}

	isExcludedPath := isAuthExcludedPath(ctx)
	isAnonymousPath := isAnonymousPath(ctx)

	if accessToken != nil {

		ctx.Values().Set(common.CTX_ACCESS_TOKEN_KEY, accessToken)
	}

	if isExcludedPath && accessToken != nil {

		return errors.New("bad request")
	}

	if isAnonymousPath {

		return nil
	}

	if accessToken == nil {

		return nil
	}

	err = checkAuthoritySessionUseCase.Execute(ctx)

	if err != nil {

		ctx.Values().Remove(common.CTX_ACCESS_TOKEN_KEY)

		return err
	}

	return nil
}
