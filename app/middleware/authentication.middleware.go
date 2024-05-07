package middleware

import (
	"app/app/model"
	"app/app/service"

	"github.com/kataras/iris/v12"
)

const AUTH_USER = "auth_user"

type AuthUser *model.User

func Authentication() func(iris.Context, *service.AuthenticateService) {

	return func(ctx iris.Context, service *service.AuthenticateService) {

		var reqToken string = ctx.GetHeader("bearer")

		if reqToken == "" {

			noToken(ctx)
			return
		}

		ctx.RegisterDependency(user)
		ctx.Values().Set(AUTH_USER, user)
		ctx.Next()
	}
}

func handleInvalidInput(ctx iris.Context, err error) {

}

func noToken(ctx iris.Context) {

	ctx.StatusCode(401)

	ctx.JSON(struct {
		Message string `json:"message"`
	}{
		Message: "unauthenticated",
	})

	ctx.EndRequest()
}
