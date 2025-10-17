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
	update struct {
		auth.Authorize[constraint.TenantAgent]
		crud.UpdateEndpointCurator
		common.Controller
	}
)

func (this *update) ENDPOINT_ModifyAssignment(
	action.UseCaseFor[requestPresenter.ModifyAssignment, responsePresenter.ModifyAssignment],
) endpoint.IEndpoint {

	return this.PATCH("/{assignmentUUID:uuid}").BuildAction(nil)
}
