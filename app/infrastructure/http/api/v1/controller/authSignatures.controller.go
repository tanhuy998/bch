package controller

import (
	"app/infrastructure/http/common"
	routeAuth "app/infrastructure/http/common/auth/route"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthSignaturesController struct {
		common.Controller
		RetrieveTenantSignaturesUseCase usecasePort.IUseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
		RefreshTenantSignaturesUseCase  usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
		RevokeTenantSignaturesUseCase   usecasePort.IUseCase[requestPresenter.Logout, responsePresenter.Logout]
	}
)

func (this *AuthSignaturesController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	routeAuth.Exclude(
		activator.Handle(
			"GET", "/tenant/{tenantUUID:uuid}", "RetrieveTenantSignatures",
			middleware.BindRequest[requestPresenter.SwitchTenant](
				container,
			),
		),
	)

	routeAuth.Exclude(
		activator.Handle(
			"POST", "/", "RefreshTenantSignatures",
			middleware.BindRequest[requestPresenter.RefreshLoginRequest](container),
		),
	)

	routeAuth.Exclude(
		activator.Handle(
			"DELETE", "/", "RevokeTenantSignatures",
			middleware.BindRequest[requestPresenter.Logout](
				container,
				middlewareHelper.UseTenantMapping,
			),
		),
	)
}

func (this *AuthSignaturesController) RetrieveTenantSignatures(
	input *requestPresenter.SwitchTenant,
) (mvc.Result, error) {

	return this.ResultOf(
		this.RetrieveTenantSignaturesUseCase.Execute(input),
	)
}

func (this *AuthSignaturesController) RefreshTenantSignatures(
	input *requestPresenter.RefreshLoginRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.RefreshTenantSignaturesUseCase.Execute(input),
	)
}

func (this *AuthSignaturesController) RevokeTenantSignatures(
	input *requestPresenter.Logout,
) (mvc.Result, error) {

	return this.ResultOf(
		this.RevokeTenantSignaturesUseCase.Execute(input),
	)
}
