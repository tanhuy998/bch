package middleware

import (
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware/middlewareHelper"
	accessTokenServicePort "app/port/accessToken"
	accessTokenClientPort "app/port/accessTokenClient"
	jwtTokenServicePort "app/port/jwtTokenService"
	"errors"
	"fmt"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
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

type (
	SigningMethod jwt.SigningMethodECDSA
	// 	Message string `json:"message"`
	// }
)

// func Authentication() func(iris.Context, authService.IAuthenticate) {

// 	return func(ctx iris.Context, auth authService.IAuthenticate) {

// 		ctx.Next()
// 	}
// }

func Auth(
	container *hero.Container,
	constraints ...middlewareHelper.AuthorityConstraint,
) iris.Handler {

	return func(ctx iris.Context) {

		accessToken := common.GetAccessToken(ctx)

		if accessToken == nil {

			container.Handler(authentication_func)(ctx)
		}

		accessToken = common.GetAccessToken(ctx)

		// ctx.Err() check whether the parent context of iris.Context was done
		// when the authorization progress failed
		if accessToken == nil {

			return
		}

		if len(constraints) == 0 {

			ctx.Next()
			return
		}

		//container.Handler(authorization_func)(ctx)

		if !validateAuthority(accessToken, constraints) {

			common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusForbidden, "forbiden authority")
			return
		}

		ctx.Next()
	}
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
This func argument is injected by the dependency injection container
*/
func authentication_func(
	ctx iris.Context,
	accessTokenClient accessTokenClientPort.IAccessTokenClient,
	accessTokenManipulator accessTokenServicePort.IAccessTokenManipulator,
) {

	accessToken, err := accessTokenClient.Read(ctx)

	//accessToken, err := accessTokenManipulator.Read(tokenString)

	switch err {
	case nil:
	case jwtTokenServicePort.ERR_SIGNING_METHOD_MISMATCH:
		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, err.Error())
	default:
		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusUnauthorized, "unauthorized")
	}

	if err != nil {

		return
	}

	errCode, err := validateAccessToken(accessToken)

	if err != nil {

		common.SendDefaulJsonBodyAndEndRequest(ctx, errCode, err.Error())
		return
	}

	//ctx.RegisterDependency(accessToken)
	ctx.Values().Set(common.CTX_ACCESS_TOKEN_KEY, accessToken)
	//ctx.Next()
}

func validateAccessToken(accessToken accessTokenServicePort.IAccessToken) (errorCode int, err error) {

	switch {
	case accessToken == nil:
		return http.StatusUnauthorized, fmt.Errorf("unauthorized")
	case accessToken.Expired():
		return http.StatusUnauthorized, ERR_ACCESS_TOKEN_EXPIRED
	default:
		return 0, nil
	}
}
