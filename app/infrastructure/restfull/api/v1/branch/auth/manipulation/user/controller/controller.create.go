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
	create struct {
		common.Controller
		crud.CreateEndpointCurator
		CreateUserUsecase usecasePort.IUseCase[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter] // usecase.ICreateUser

	}
)

func (this *create) ENDPOINT_CreateUser(
	input.Bind[requestPresenter.CreateUserRequestPresenter],
) endpoint.IEndpoint {
	return this.POST("/").
		BuildAction(
			func(input *requestPresenter.CreateUserRequestPresenter) (mvc.Result, error) {

				return this.ResultOf(
					this.CreateUserUsecase.Execute(input),
				)
			},
		)
}
