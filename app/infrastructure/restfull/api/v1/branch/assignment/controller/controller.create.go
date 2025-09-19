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
		crud.CreateEndpointCurator
		common.Controller
		CreateAssignmentUseCase                     usecasePort.IUseCase[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse]
		CreateAssignmentGroupUseCase                usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse]
		AddCommandGroupUserToAssignmentGroupUseCase usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupMember, responsePresenter.CreateAssignmentGroupMemeber]
	}
)

func (this *create) ENDPOINT_CreateAssignment(
	auth.Authorize[constraint.TenantAgent],
	input.Bind[requestPresenter.CreateAssigmentRequest],
) endpoint.IEndpoint {

	return this.POST("/").
		BuildAction(
			func(input *requestPresenter.CreateAssigmentRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.CreateAssignmentUseCase.Execute(input),
				)
			},
		)
}

func (this *create) ENDPOINT_CreateAssignmentGroup(
	auth.Authorize[constraint.TenantAgent],
	input.Bind[requestPresenter.CreateAssignmentGroupRequest],
) endpoint.IEndpoint {

	return this.POST("/{assignmentUUID:uuid}/group/command/{commandGroupUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.CreateAssignmentGroupRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.CreateAssignmentGroupUseCase.Execute(input),
				)
			},
		)
}

func (this *create) ENDPOINT_CreateAssignmentGroupMember(
	auth.Authorize[constraint.TenantAgent],
	input.Bind[requestPresenter.CreateAssignmentGroupMember],
) endpoint.IEndpoint {

	return this.POST("/group/{groupUUID:uuid}/member").
		BuildAction(
			func(input *requestPresenter.CreateAssignmentGroupMember) (mvc.Result, error) {
				return this.ResultOf(
					this.AddCommandGroupUserToAssignmentGroupUseCase.Execute(input),
				)
			},
		)
}
