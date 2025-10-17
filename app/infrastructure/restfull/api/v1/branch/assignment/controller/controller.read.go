package controller

import (
	"app/infrastructure/restfull/common/annotation/action"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	"app/model"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	read struct {
		crud.ReadEndpointCurator
	}
)

func (this *read) ENDPOINT_GetAssignments(
	auth.Authorize[constraint.TenantAgent],
	action.UseCaseFor[requestPresenter.GetAssignments, responsePresenter.GetAssignments[model.Assignment]],
) endpoint.IEndpoint {

	return this.GET("/").BuildAction(nil)
}

func (this *read) ENDPOINT_GetSingleAssignment(
	action.UseCaseFor[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse],
) endpoint.IEndpoint {

	return this.GET("/{uuid:uuid}").BuildAction(nil)
}

func (this *read) ENDPOINT_GetAssignmentGroups(
	auth.Authorize[constraint.TenantAgent],
	action.UseCaseFor[requestPresenter.GetAssignmentGroups, responsePresenter.GetAssignmentGroups[model.AssignmentGroup]],
) endpoint.IEndpoint {
	return this.GET("/{uuid:uuid}/group/list").BuildAction(nil)
}
