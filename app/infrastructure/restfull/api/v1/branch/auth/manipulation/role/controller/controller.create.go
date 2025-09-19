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
		GrantCommandGroupRolesToUserUseCase usecasePort.IUseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse]
	}
)

func (this *create) ENDPOINT_GrantCommandGroupRolesToUser(
	input.Bind[requestPresenter.GrantCommandGroupRolesToUserRequest],
) endpoint.IEndpoint {

	return this.POST("/group/{groupUUID}/user/{userUUID}").
		BuildAction(
			func(input *requestPresenter.GrantCommandGroupRolesToUserRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.GrantCommandGroupRolesToUserUseCase.Execute(input),
				)
			},
		)
}
