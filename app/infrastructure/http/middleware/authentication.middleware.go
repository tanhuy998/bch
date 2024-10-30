package middleware

import (
	"app/infrastructure/http/common"
	libAuth "app/infrastructure/http/common/auth/lib"
	"app/infrastructure/http/middleware/middlewareHelper"
	internalCommon "app/internal/common"
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

	return func(ctx iris.Context) {

		accessToken := common.GetAccessToken(ctx)

		if accessToken == nil {
			// do dependencies injection
			container.Handler(authentication_func)(ctx)

		}

		if ctx.IsStopped() {
			/**
			authentication will stop the excution of the handler chain when
			the request context doesn't match the route's requirements
			*/
			return
		}

		accessToken = common.GetAccessToken(ctx)

		if accessToken == nil {
			/*
				No access token but but the reqeust context was not stopped
				means the current request context's path is auth excluded
			*/
			ctx.Next()
			return
		}

		if len(constraints) == 0 {

			ctx.Next()
			return
		}

		if !validateAuthority(accessToken, constraints) {

			common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusForbidden, "forbiden authority")
			return
		}

		ctx.Next()
	}
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
) {

	accessToken, err := accessTokenClient.Read(ctx)

	switch {
	case err == nil:
	// case jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH:
	// 	common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, err.Error())
	case errors.Is(err, internalCommon.ERR_INTERNAL):
		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, "internal error")
	default:
		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusUnauthorized, "(Authentication middleware error) unauthorized")
	}

	if err != nil {

		return
	}

	isExcludedPath := isAuthExcludedPath(ctx)
	isAnonymousPath := isAnonymousPath(ctx)

	if accessToken != nil {

		ctx.Values().Set(common.CTX_ACCESS_TOKEN_KEY, accessToken)
	}

	if isExcludedPath && accessToken != nil {

		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusBadRequest, "bad request")
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

		if errors.Is(err, internalCommon.ERR_INTERNAL) {

			common.SendDefaulJsonBodyAndEndRequest(ctx, 500, err.Error())
			return
		}

		common.SendDefaulJsonBodyAndEndRequest(ctx, 401, err.Error())
		return
	}
}
