package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	create struct {
		crud.CreateEndpointCurator
		common.Controller
		CreateTenantUseCase usecasePort.IUseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse]
	}
)

func (this *create) ENDPOINT_CreateTenant(
	input.Bind[requestPresenter.CreateTenantRequest],
) endpoint.IEndpoint {
	return this.POST("/").
		BuildAction(
			func(input *requestPresenter.CreateTenantRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.CreateTenantUseCase.Execute(input),
				)
			},
		)
}
