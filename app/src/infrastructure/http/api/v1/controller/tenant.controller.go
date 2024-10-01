package controller

import (
	"app/src/infrastructure/http/middleware"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	usecase "app/src/useCase"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type (
	TenantController struct {
		// CreateTenantAgentUsecase usecase.ICreateTenantAgent
		CreateTenantUseCase usecase.ICreateTenant
	}
)

func (this *TenantController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	activator.Handle(
		"POST", "/", "CreateTenant",
		middleware.BindPresenters[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse](container),
	)
}

// func (this *TenantController) CreateTenantAgent(
// 	input *requestPresenter.CreateTenantAgentRequest,
// 	output *responsePresenter.CreateTenantAgentResponse,
// ) (mvc.Result, error) {

// 	return this.CreateTenantAgentUsecase.Execute(input, output)
// }

func (this *TenantController) CreateTenant(
	input *requestPresenter.CreateTenantRequest,
	output *responsePresenter.CreateTenantResponse,
	ctx iris.Context,
) (mvc.Result, error) {

	return this.CreateTenantUseCase.Execute(input, output)
}
