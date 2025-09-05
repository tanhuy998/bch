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
	update struct {
		crud.UpdateEndpointCurator
		common.Controller
		ModifyAssignmentUseCase usecasePort.IUseCase[requestPresenter.ModifyAssignment, responsePresenter.ModifyAssignment]
	}
)

func (this *update) ENDPOINT_ModifyAssignment() endpoint.IEndpoint {

	return this.PATCH("/{assignmentUUID:uuid}").
		UseMiddleware(
			middleware.Auth(
				hook.AuthRequireTenantAgent,
			),
			middleware.BindRequest[requestPresenter.ModifyAssignment](
				hook.UseAuthority,
				hook.UseTenantMapping,
			),
		).
		Build()
}
func (this *update) ModifyAssignment(
	input *requestPresenter.ModifyAssignment,
) (mvc.Result, error) {

	return this.ResultOf(
		this.ModifyAssignmentUseCase.Execute(input),
	)
}
