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
		GetParitcipatedCommandGroupUseCase              usecasePort.IUseCase[requestPresenter.GetUserParticipatedCommandGroups, responsePresenter.GetUserParticipatedCommandGroups[model.CommandGroup]]
		GetTenantCommandGroupsUseCase                   usecasePort.IUseCase[requestPresenter.GetTenantCommandGroups, responsePresenter.GetTenantCommandGroups[model.CommandGroup]]
		GetAssignmentUnAssignedCommandGroupUsersUseCase usecasePort.IUseCase[requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers, responsePresenter.GetAssignmentGroupUnAssignedCommandGroupUsers[model.CommandGroupUser]]
	}
)

func (this *read) ENDPOINT_GetParticipatedGroups(
	input.Bind[requestPresenter.GetUserParticipatedCommandGroups],
) endpoint.IEndpoint {
	return this.GET("/participated/user/{userUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.GetUserParticipatedCommandGroups) (mvc.Result, error) {

				return this.ResultOf(
					this.GetParitcipatedCommandGroupUseCase.Execute(input),
				)
			},
		)
}

func (this *read) ENDPOINT_GetCommandGroups(
	input.Bind[requestPresenter.GetTenantCommandGroups],
) endpoint.IEndpoint {
	return this.GET("/").
		BuildAction(
			func(input *requestPresenter.GetTenantCommandGroups) (mvc.Result, error) {

				return this.ResultOf(
					this.GetTenantCommandGroupsUseCase.Execute(input),
				)
			},
		)
}

func (this *read) ENDPOINT_GetUnAssignedCommandGroupUsers(
	input.Bind[requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers],
) endpoint.IEndpoint {
	return this.GET("/users/unassigned/{assignmentGroupUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers) (mvc.Result, error) {

				return this.ResultOf(
					this.GetAssignmentUnAssignedCommandGroupUsersUseCase.Execute(input),
				)
			},
		)
}
