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
	update struct {
		crud.UpdateEndpointCurator
		common.Controller
		ModifyAssignmentUseCase usecasePort.IUseCase[requestPresenter.ModifyAssignment, responsePresenter.ModifyAssignment]
	}
)

func (this *update) ENDPOINT_ModifyAssignment(
	auth.Authorize[constraint.TenantAgent],
	input.Bind[requestPresenter.ModifyAssignment],
) endpoint.IEndpoint {

	return this.PATCH("/{assignmentUUID:uuid}").
		BuildAction(
			func(input *requestPresenter.ModifyAssignment) (mvc.Result, error) {
				return this.ResultOf(
					this.ModifyAssignmentUseCase.Execute(input),
				)
			},
		)
}
