package controller

import (
	loginDomain "app/domain/auth/login"
	refreshLoginDomain "app/domain/auth/refreshLogin"
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
	libConfig "app/internal/lib/config"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type (
	UserLoggingController struct {
		common.Controller
		LoginUseCase        usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse]
		RefreshLoginUseCase usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
	}
)

func (this *UserLoggingController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	this.bindDependencies(container)

	activator.Handle(
		"POST", "/login", "LogIn",
		middleware.BindPresenters[requestPresenter.LoginRequest, responsePresenter.LoginResponse](container),
	)

	activator.Handle(
		"POST", "/refresh", "Refresh",
		middleware.BindPresenters[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse](container),
	)
}

func (this *UserLoggingController) bindDependencies(container *hero.Container) {

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse],
		loginDomain.LogInUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse],
		refreshLoginDomain.RefreshLoginUseCase,
	](container, nil)
}

func (this *UserLoggingController) LogIn(
	input *requestPresenter.LoginRequest,
	output *responsePresenter.LoginResponse,
) (mvc.Result, error) {

	return this.ResultOf(
		this.LoginUseCase.Execute(input),
	)
}

func (this *UserLoggingController) Refresh(
	input *requestPresenter.RefreshLoginRequest,
	output *responsePresenter.RefreshLoginResponse,
) (mvc.Result, error) {

	return this.ResultOf(
		this.RefreshLoginUseCase.Execute(input),
	)
}

func (this *UserLoggingController) Temp() {

}
