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
	TenantController struct {
		// CreateTenantAgentUsecase usecase.ICreateTenantAgent
		Controller          *common.Controller
		CreateTenantUseCase usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse]
	}
)

func (this *TenantController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	activator.Handle(
		"POST", "/", "CreateTenant",
		middleware.BindRequest[requestPresenter.CreateTenantRequest](container),
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
