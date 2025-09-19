package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	update struct {
		common.Controller
		crud.UpdateEndpointCurator
		ModifyUserUsecase usecasePort.IUseCase[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse] // usecase.IModifyUser

	}
)

func (this *update) ENDPOINT_ModifyUser(
	input.Bind[requestPresenter.ModifyUserRequest],
) endpoint.IEndpoint {
	return this.PATCH("/{userUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.ModifyUserRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.ModifyUserUsecase.Execute(input),
				)
			},
		)
}
