package middleware

import (
	authService "app/app/service/auth"
	"fmt"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12"
)

const (
	AUTH_USER      = "auth_user"
	AUTH_HEADER    = "AUTH_HEADER"
	JWT_PUBLIC_KEY = "JWT_PUBLIC_KEY"
)

type SigningMethod jwt.SigningMethodECDSA

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

func handleInvalidInput(ctx iris.Context, err error) {

}

func unAuthenticated(ctx iris.Context) {

	ctx.StatusCode(401)

	ctx.JSON(struct {
		Message string `json:"message"`
	}{
		Message: "unauthenticated",
	})

	ctx.EndRequest()
}

func verifyToken(strToken string) (*jwt.Token, error) {

	// token := jwt.New(jwt.SigningMethodES256)
	// token.Raw = strToken

	return jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {

			return nil, fmt.Errorf("")
		}

		return JWT_PUBLIC_KEY, nil
	})
}
