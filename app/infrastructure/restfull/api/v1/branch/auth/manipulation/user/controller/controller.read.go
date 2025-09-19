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
		GetGroupUserUsecase   usecasePort.IUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse] // usecase.IGetGroupUsers
		ModifyUserUsecase     usecasePort.IUseCase[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse]       // usecase.IModifyUser
		GetTenantUsersUseCase usecasePort.IUseCase[requestPresenter.GetTenantUsers, responsePresenter.GetTenantUsers[model.User]]
	}
)

func (this *read) ENDPOINT_GetGroupUsers(
	input.Bind[requestPresenter.GetGroupUsersRequest],
) endpoint.IEndpoint {
	return this.GET("/group/{groupUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.GetGroupUsersRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.GetGroupUserUsecase.Execute(input),
				)
			},
		)
}

func (this *read) ENDPOINT_GetTenantUsers(
	input.Bind[requestPresenter.GetTenantUsers],
) endpoint.IEndpoint {
	return this.GET("/").
		BuildAction(
			func(input *requestPresenter.GetTenantUsers) (mvc.Result, error) {

				return this.ResultOf(
					this.GetTenantUsersUseCase.Execute(input),
				)
			},
		)
}
