package controller

import (
	"app/infrastructure/http/common"
	routeAuth "app/infrastructure/http/common/auth/route"
	"app/infrastructure/http/middleware"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthGeneralController struct {
		common.Controller
		AuthenticateCredentialsUseCase usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse]
		NavigateTenantUseCase          usecasePort.IUseCase[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant]
		CheckGenTokenUseCase           usecasePort.IUseCase[requestPresenter.CheckLogin, responsePresenter.CheckLogin]
	}
)

func (this *AuthGeneralController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	routeAuth.Exclude(
		activator.Handle(
			"POST", "/credentials", "AuthenticateCredentials",
			middleware.AuthPolicies(container),
			middleware.BindRequest[requestPresenter.LoginRequest](container),
		),
	)

	routeAuth.Exclude(
		activator.Handle(
			"GET", "/nav", "NavigateTenant",
			middleware.BindRequest[requestPresenter.AuthNavigateTenant](container),
		),
	)

	routeAuth.Exclude(
		activator.Handle(
			"HEAD", "/", "CheckGenToken",
			middleware.BindRequest[requestPresenter.CheckLogin](container),
		),
	)
}

func (this *AuthGeneralController) AuthenticateCredentials(
	input *requestPresenter.LoginRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.AuthenticateCredentialsUseCase.Execute(input),
	)
}

func (this *AuthGeneralController) NavigateTenant(
	input *requestPresenter.AuthNavigateTenant,
) (mvc.Result, error) {

	return this.ResultOf(
		this.NavigateTenantUseCase.Execute(input),
	)
}

func (this *AuthGeneralController) CheckGenToken(
	input *requestPresenter.CheckLogin,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CheckGenTokenUseCase.Execute(input),
	)
}
