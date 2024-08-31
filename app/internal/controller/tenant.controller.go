package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	TenantController struct {
		CreateTenantAgentUsecase usecase.ICreateTenantAgent
	}
)

func (this *TenantController) CreateTenantAgent(
	input *requestPresenter.CreateTenantAgentRequest,
	output *responsePresenter.CreateTenantAgentResponse,
) (mvc.Result, error) {

	return this.CreateTenantAgentUsecase.Execute(input, output)
}

func (this *TenantController) CreateTenant() {

}
