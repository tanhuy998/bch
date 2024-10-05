package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
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

	container := activator.Dependencies()

	activator.Handle(
		"POST", "/login", "LogIn",
		middleware.BindPresenters[requestPresenter.LoginRequest, responsePresenter.LoginResponse](container),
	)

	activator.Handle(
		"POST", "/refresh", "Refresh",
		middleware.BindPresenters[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse](container),
	)
}

func (this *UserLoggingController) BindDependencies(container *hero.Container) common.IController {

	return this
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
