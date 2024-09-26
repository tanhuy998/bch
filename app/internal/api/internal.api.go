package api

import (
	accessTokenServicePort "app/adapter/accessToken"
	"app/internal/api/auth/userLogging"
	"app/internal/middleware"
	libConfig "app/lib/config"
	"app/service/accessTokenService"

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
