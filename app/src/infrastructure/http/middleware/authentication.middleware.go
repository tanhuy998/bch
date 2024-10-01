package middleware

import (
	"app/src/infrastructure/http/common"
	"app/src/infrastructure/http/middleware/middlewareHelper"
	accessTokenServicePort "app/src/port/accessToken"
	accessTokenClientPort "app/src/port/accessTokenClient"
	jwtTokenServicePort "app/src/port/jwtTokenService"
	"errors"
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

	// return container.Handler(
	// 	authentication_func,
	// )

	// to ensure constraint function list not be changed when
	// server is running
	copy(constraints, constraints[:])

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

		if !validateAuthority(accessToken.GetAuthData(), constraints) {

			common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusForbidden, "user authority unauthorized")
			return
		}

		ctx.Next()
	}
}

// func authorization_func() (passed bool) {

// 	return true
// }

func validateAuthority(authority accessTokenServicePort.IAccessTokenAuthData, constraints []middlewareHelper.AuthorityConstraint) bool {

	if authority == nil {

		return false
	}

	for _, f := range constraints {

		if !f(authority) {

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

	tokenString := accessTokenClient.Read(ctx)

	accessToken, err := accessTokenManipulator.Read(tokenString)

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
	case accessToken.Expired():
		return http.StatusUnauthorized, ERR_ACCESS_TOKEN_EXPIRED
	default:
		return 0, nil
	}
}
