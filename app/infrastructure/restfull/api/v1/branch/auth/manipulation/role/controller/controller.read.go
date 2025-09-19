package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	"app/model"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	read struct {
		common.Controller
		crud.ReadEndpointCurator
		GetAllRolesUseCase usecasePort.IUseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse[model.Role]]
	}
)

func (this *read) ENDPOINT_GetAllRoles(
	input.Bind[requestPresenter.GetAllRolesRequest],
) endpoint.IEndpoint {
	return this.GET("/").
		BuildAction(
			func(input *requestPresenter.GetAllRolesRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.GetAllRolesUseCase.Execute(input),
				)
			},
		)
}
