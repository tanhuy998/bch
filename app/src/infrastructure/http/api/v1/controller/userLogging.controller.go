package controller

import (
	"app/src/infrastructure/http/middleware"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	usecase "app/src/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	UserLoggingController struct {
		LoginUseCase        usecase.ILogIn
		RefreshLoginUseCase usecase.IRefreshLogin
	}
)

func (this *UserLoggingController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	activator.Handle(
		"POST", "/login", "LogIn",
		middleware.BindPresenters[requestPresenter.LoginRequest, responsePresenter.LoginResponse](container),
	)

	activator.Handle(
		"POST", "/refresh", "Refresh",
		middleware.BindPresenters[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse](container),
	)
}

func (this *UserLoggingController) LogIn(
	input *requestPresenter.LoginRequest,
	output *responsePresenter.LoginResponse,
) (mvc.Result, error) {

	return this.LoginUseCase.Execute(input, output)
}

func (this *UserLoggingController) Refresh(
	input *requestPresenter.RefreshLoginRequest,
	output *responsePresenter.RefreshLoginResponse,
) (mvc.Result, error) {

	return this.RefreshLoginUseCase.Execute(input, output)
}

func (this *UserLoggingController) Temp() {

}
