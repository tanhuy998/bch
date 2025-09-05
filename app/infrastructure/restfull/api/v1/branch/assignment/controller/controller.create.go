package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	"app/infrastructure/restfull/common/middleware"
	"app/infrastructure/restfull/common/middleware/hook"
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

func (this *create) ENDPOINT_CreateAssignment() endpoint.IEndpoint {

	return this.POST("/").
		UseMiddleware(
			middleware.Auth(
				hook.AuthRequireTenantAgent,
			),
			middleware.BindRequest[requestPresenter.CreateAssigmentRequest](
				hook.UseAuthority,
				hook.UseTenantMapping,
			),
		).
		Build()
}
func (this *create) CreateAssignment(
	input *requestPresenter.CreateAssigmentRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CreateAssignmentUseCase.Execute(input),
	)
}

func (this *create) ENDPOINT_CreateAssignmentGroup() endpoint.IEndpoint {

	return this.POST("/{assignmentUUID:uuid}/group/command/{commandGroupUUID:uuid}").
		UseMiddleware(
			middleware.Auth(
				hook.AuthRequireTenantAgent,
			),
			middleware.BindRequest[requestPresenter.CreateAssignmentGroupRequest](
				hook.UseAuthority,
				hook.UseTenantMapping,
			),
		).
		Build()
}
func (this *create) CreateAssignmentGroup(
	input *requestPresenter.CreateAssignmentGroupRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CreateAssignmentGroupUseCase.Execute(input),
	)
}

func (this *create) ENDPOINT_CreateAssignmentGroupMember() endpoint.IEndpoint {

	return this.POST("/group/{groupUUID:uuid}/member").
		UseMiddleware(
			middleware.Auth(
				hook.AuthRequiredTenantAgentExceptMeetRoles("COMMANDER"),
			),
			middleware.BindRequest[requestPresenter.CreateAssignmentGroupMember](
				hook.UseAuthority,
				hook.UseTenantMapping,
			),
		).
		Build()
}
func (this *create) CreateAssignmentGroupMember(
	input *requestPresenter.CreateAssignmentGroupMember,
) (mvc.Result, error) {

	return this.ResultOf(
		this.AddCommandGroupUserToAssignmentGroupUseCase.Execute(input),
	)
}
