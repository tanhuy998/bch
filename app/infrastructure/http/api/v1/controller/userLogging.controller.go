package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type (
	UserLoggingController struct {
		common.Controller
		LoginUseCase          usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse]
		RefreshLoginUseCase   usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
		NavigateTenantUseCase usecasePort.IUseCase[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant]
		LogoutUseCase         usecasePort.IUseCase[requestPresenter.Logout, responsePresenter.Logout]
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

	activator.Handle(
		"GET", "/nav", "NavigateTenant",
		middleware.BindRequest[requestPresenter.AuthNavigateTenant](container),
	)

	activator.Handle(
		"DELETE", "/logout", "Logout",
		middleware.BindRequest[requestPresenter.Logout](
			container,
			middlewareHelper.UseTenantMapping,
		),
	)
}

func (this *UserLoggingController) BindDependencies(container *hero.Container) common.IController {

	return this
}

func (this *UserLoggingController) LogIn(
	input *requestPresenter.LoginRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.LoginUseCase.Execute(input),
	)
}

func (this *UserLoggingController) Refresh(
	input *requestPresenter.RefreshLoginRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.RefreshLoginUseCase.Execute(input),
	)
}

func (this *UserLoggingController) Temp() {

}

func (this *UserLoggingController) NavigateTenant(
	input *requestPresenter.AuthNavigateTenant,
) (mvc.Result, error) {

	return this.ResultOf(
		this.NavigateTenantUseCase.Execute(input),
	)
}

func (this *UserLoggingController) Logout(
	input *requestPresenter.Logout,
) (mvc.Result, error) {

	return this.ResultOf(
		this.LogoutUseCase.Execute(input),
	)
}
