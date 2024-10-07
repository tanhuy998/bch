package api

import (
	loginDomain "app/domain/auth/login"
	"app/infrastructure/http/api/v1/branch/auth/userLogging"
	"app/infrastructure/http/middleware"
	libConfig "app/internal/lib/config"
	accessTokenServicePort "app/port/accessToken"
	authServicePort "app/port/auth"
	authSignatureTokenPort "app/port/authSignatureToken"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/service/accessTokenService"
	"app/service/authSignatureToken"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
)

func initInternalAPI(app router.Party) {

	router := app.Party("/internal", middleware.SecretAuth)

	authRouter := router.Party("/auth")

	authRouter.ConfigureContainer().Use(
		func(ctx iris.Context, accessTokenManipulator accessTokenServicePort.IAccessTokenManipulator) {

			ctx.Values().Set(accessTokenManipulator.CtxNoExpireKey(), true)
			ctx.Next()
		},
	)

	userLogging.RegisterUserLoggingApi(authRouter).EnableStructDependents()
	initTenantApi(router).EnableStructDependents()
}

func bindDependencies(container *hero.Container) {

	libConfig.OverrideDependency[accessTokenServicePort.IAccessTokenManipulator](
		container, accessTokenService.New(accessTokenService.WithoutExpire),
	)
	libConfig.OverrideDependency[authSignatureTokenPort.IAuthSignatureProvider, authSignatureToken.AuthSignatureTokenService](container, nil)
	libConfig.OverrideDependency[authServicePort.ILogIn, loginDomain.LogInService](container, nil)
	libConfig.OverrideDependency[
		usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse],
		loginDomain.LogInUseCase,
	](container, nil)
}
