package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
	accessTokenClientPort "app/port/accessTokenClient"
	tenantAccessTokenServicePort "app/port/generalTokenServicePort"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type (
	TenantController struct {
		// CreateTenantAgentUsecase usecase.ICreateTenantAgent
		Controller                    *common.Controller
		AccessTokenClient             accessTokenClientPort.IAccessTokenClient
		TenantAccessTokenManipulaotr  tenantAccessTokenServicePort.IGeneralTokenManipulator
		CreateTenantUseCase           usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse]
		GrantUserAsTenantAgentUseCase usecasePort.IUseCase[requestPresenter.GrantUserAsTenantAgent, responsePresenter.GrantUserAsTenantAgent]
		SwitchTenantUseCase           usecasePort.IUseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
	}
)

func (this *TenantController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	activator.Handle(
		"POST", "/", "CreateTenant",
		middleware.BindRequest[requestPresenter.CreateTenantRequest](container),
	)

	activator.Handle(
		"GET", "/agent/grant/{userUUID:uuid}", "GrantUserAsTenantAgent",
		middleware.BindRequest[requestPresenter.CreateTenantRequest](container),
	)

	activator.Handle(
		"GET", "/switch", "SwitchTenant",
		middleware.BindRequest[requestPresenter.SwitchTenant](container),
	)
}

func (this *TenantController) BindDependencies(container *hero.Container) common.IController {

	return this
}

// func (this *TenantController) CreateTenantAgent(
// 	input *requestPresenter.CreateTenantAgentRequest,
// 	output *responsePresenter.CreateTenantAgentResponse,
// ) (mvc.Result, error) {

// 	return this.CreateTenantAgentUsecase.Execute(input, output)
// }

func (this *TenantController) CreateTenant(
	input *requestPresenter.CreateTenantRequest,
) (mvc.Result, error) {

	return this.Controller.ResultOf(
		this.CreateTenantUseCase.Execute(input),
	)
}

func (this *TenantController) GrantUserAsTenantAgent(
	input *requestPresenter.GrantUserAsTenantAgent,
) (mvc.Result, error) {

	return this.Controller.ResultOf(
		this.GrantUserAsTenantAgentUseCase.Execute(input),
	)
}

func (this *TenantController) SwitchTenant(
	input *requestPresenter.SwitchTenant,
) (mvc.Result, error) {

	return this.Controller.ResultOf(
		this.SwitchTenantUseCase.Execute(input),
	)
}
