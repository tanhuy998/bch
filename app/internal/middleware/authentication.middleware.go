package middleware

import (
	accessTokenServicePort "app/adapter/accessToken"
	authServiceAdapter "app/adapter/auth"
	"app/internal"
	"app/internal/common"
	authService "app/service/auth"
	jwtTokenService "app/service/jwtToken"
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

const (
	AUTH_USER                = "auth_user"
	AUTH_HEADER              = "AUTH_HEADER"
	JWT_PUBLIC_KEY           = "JWT_PUBLIC_KEY"
	AUTH_REQ_HEADER          = "Authorization"
	AUTH_REQ_HEADER_SCHEME   = "bearer "
	AUTH_COOKIE_ACCESS_TOKEN = "access-token"
)

var (
	ERR_INVALID_ACCESS_TOKEN = errors.New("invalid access token")
	ERR_MISSING_AUTH_HEADER  = errors.New("missing authorization header")
	ERR_ACCESS_TOKEN_EXPIRED = errors.New("access token expired")
)

type (
	SigningMethod jwt.SigningMethodECDSA

	// auth_err_body_reponse struct {
	// 	Message string `json:"message"`
	// }
)

func Authentication() func(iris.Context, authService.IAuthenticate) {

	return func(ctx iris.Context, auth authService.IAuthenticate) {
		fmt.Println("1")
		// ENV_AUTH_HEADER := env.Get(AUTH_HEADER, "bearer")
		// var strToken string = ctx.GetHeader(ENV_AUTH_HEADER)

		// if strToken == "" {

		// 	unAuthenticated(ctx)
		// 	return
		// }

		// token, err := verifyToken(strToken)

		// if err != nil {

		// 	unAuthenticated(ctx)
		// 	return
		// }

		// ctx.RegisterDependency(token)
		//ctx.Values().Set(AUTH_USER, user)
		ctx.Next()
	}
}

func AuthenticationBearer(container *hero.Container) iris.Handler {

	return container.Handler(
		fetch_and_store_user_data,
	)
}

func Auth(container *hero.Container) iris.Handler {

	return container.Handler(
		authentication_func,
	)
}

func authentication_func(ctx iris.Context, accessTokenHandler accessTokenServicePort.IAccessTokenManipulator) {

	tokenString, err := retrieveTokenString(ctx)

	if err != nil {

		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := accessTokenHandler.Read(tokenString)

	switch err {
	case nil:
	case jwtTokenService.ERR_SIGNING_METHOD_MISMATCH:
		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusInternalServerError, err.Error())
	default:
		common.SendDefaulJsonBodyAndEndRequest(ctx, http.StatusUnauthorized, err.Error())
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
	ctx.Values().Set(internal.CTX_ACCESS_TOKEN_KEY, accessToken)
	ctx.Next()
}

func validateAccessToken(accessToken accessTokenServicePort.IAccessToken) (errorCode int, err error) {

	switch {
	case accessToken.Expired():
		return http.StatusForbidden, ERR_ACCESS_TOKEN_EXPIRED
	default:
		return 0, nil
	}
}

func fetch_and_store_user_data(ctx iris.Context, accessToken *jwt.Token, fetchAuthDataService authServiceAdapter.IFetchAuthData) {

	aud, err := accessToken.Claims.GetAudience()

	if err != nil {

		return
	}

	if len(aud) == 0 || len(aud) > 1 {

		return
	}

}

func retrieveTokenString(ctx iris.Context) (string, error) {

	str := ctx.Request().Header.Get(AUTH_REQ_HEADER)

	if str == "" {

		return "", ERR_MISSING_AUTH_HEADER
	}

	token_str := strings.TrimPrefix(str, AUTH_REQ_HEADER_SCHEME)

	if token_str == "" || len(token_str) < 100 {

		return "", ERR_INVALID_ACCESS_TOKEN
	}

	return token_str, nil
}
