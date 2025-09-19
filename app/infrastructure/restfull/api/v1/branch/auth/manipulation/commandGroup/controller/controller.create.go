package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
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
		AddUserToCommandGroupUseCase usecasePort.IUseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse]
		CreateCommandGroupUseCase    usecasePort.IUseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse]
	}
)

func (this *create) ANNOTATIONS_(
	auth.Authorize[constraint.TenantAgent],
) {
}

func (this *create) ENDPOINT_CreateGroup(
	input.Bind[requestPresenter.CreateCommandGroupRequest],
) endpoint.IEndpoint {
	return this.POST("/").
		BuildAction(
			func(input *requestPresenter.CreateCommandGroupRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.CreateCommandGroupUseCase.Execute(input),
				)
			},
		)
}

func (this *create) ENDPOINT_AddUserToGroup(
	input.Bind[requestPresenter.AddUserToCommandGroupRequest],
) endpoint.IEndpoint {
	return this.POST("/{groupUUID:uuid}/user/{userUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.AddUserToCommandGroupRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.AddUserToCommandGroupUseCase.Execute(input),
				)
			},
		)
}
