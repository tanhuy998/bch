package controller

import (
	createTenantDomain "app/domain/tenant/createTenant"
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
	TenantController struct {
		// CreateTenantAgentUsecase usecase.ICreateTenantAgent
		*common.Controller
		CreateTenantUseCase usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse]
	}
)

func (this *TenantController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	this.bindDependencies(container)

	activator.Handle(
		"POST", "/", "CreateTenant",
		middleware.BindRequest[requestPresenter.CreateTenantRequest](container),
	)
}

func (this *TenantController) bindDependencies(container *hero.Container) {

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse],
		createTenantDomain.CreateTenantUseCase,
	](container, nil)
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

	return this.ResultOf(
		this.CreateTenantUseCase.Execute(input),
	)
}
