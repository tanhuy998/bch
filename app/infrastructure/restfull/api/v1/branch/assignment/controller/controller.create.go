package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/action"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	create struct {
		auth.Authorize[constraint.TenantAgent]
		crud.CreateEndpointCurator
		common.Controller
	}
)

func (this *create) ENDPOINT_CreateAssignment(
	action.UseCaseFor[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse],
) endpoint.IEndpoint {

	return this.POST("/").BuildAction(nil)
}

func (this *create) ENDPOINT_CreateAssignmentGroup(
	action.UseCaseFor[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse],
) endpoint.IEndpoint {

	return this.POST("/{assignmentUUID:uuid}/group/command/{commandGroupUUID:uuid}").BuildAction(nil)
}

func (this *create) ENDPOINT_CreateAssignmentGroupMember(
	action.UseCaseFor[requestPresenter.CreateAssignmentGroupMember, responsePresenter.CreateAssignmentGroupMemeber],
) endpoint.IEndpoint {

	return this.POST("/group/{groupUUID:uuid}/member").BuildAction(nil)
}
