package api

import (
	"app/src/infrastructure/http/api/v1/branch/auth/userLogging"
	"app/src/infrastructure/http/middleware"
	libConfig "app/src/internal/lib/config"
	accessTokenServicePort "app/src/port/accessToken"
	"app/src/service/accessTokenService"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

func initInternalAPI(app router.Party) *mvc.Application {

	router := app.Party("/internal", middleware.SecretAuth)
	// router.ConfigureContainer(
	// 	func(api *iris.APIContainer) {

	// 		api.Use(middleware.SecretAuth)
	// 	},
	// )

	authRouter := router.Party("/auth")

	authRouter.ConfigureContainer(
		func(api *iris.APIContainer) {

			bindDependencies(api.Container)
		},
	)

	return userLogging.RegisterUserLoggingApi(authRouter)
}

func bindDependencies(container *hero.Container) {

	libConfig.BindDependency[accessTokenServicePort.IAccessTokenManipulator](
		container, accessTokenService.New(accessTokenService.WithoutExpire),
	)
}
