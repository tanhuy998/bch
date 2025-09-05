package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/crud"
	"app/infrastructure/restfull/common/endpoint"
	"app/infrastructure/restfull/common/middleware"
	"app/infrastructure/restfull/common/middleware/hook"
	"app/model"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	read struct {
		crud.ReadEndpointCurator
		common.Controller
		GetSingleAssignmentUseCase  usecasePort.IUseCase[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse]
		GetAssignmentsUseCase       usecasePort.IUseCase[requestPresenter.GetAssignments, responsePresenter.GetAssignments[model.Assignment]]
		GetAssignmnentGroupsUseCase usecasePort.IUseCase[requestPresenter.GetAssignmentGroups, responsePresenter.GetAssignmentGroups[model.AssignmentGroup]]
	}
)

func (this *read) ENDPOINT_GetAssignments(
	input.Bind[requestPresenter.GetAssignments],
	input.BindInputAuthority,
	input.BindInputTenant,
) endpoint.IEndpoint {
	return this.GET("/").UseMiddleware(
		middleware.Auth(
			hook.AuthRequireTenantAgent,
		),
		// middleware.BindRequest[requestPresenter.GetAssignments](
		// 	hook.UseAuthority,
		// 	hook.UseTenantMapping,
		// ),
	).Build()
}
func (this *read) GetAssignments(
	input *requestPresenter.GetAssignments,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetAssignmentsUseCase.Execute(input),
	)
}

func (this *read) ENDPOINT_GetSingleAssignment(
	input.Bind[requestPresenter.GetSingleAssignmentRequest],
	input.BindInputAuthority,
	input.BindInputTenant,
) endpoint.IEndpoint {

	return this.GET("/{uuid:uuid}").UseMiddleware(
	// middleware.BindRequest[requestPresenter.GetSingleAssignmentRequest](
	// 	hook.UseAuthority,
	// 	hook.UseTenantMapping,
	// ),
	).Build()
}
func (this *read) GetSingleAssignment(
	input *requestPresenter.GetSingleAssignmentRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetSingleAssignmentUseCase.Execute(input),
	)
}

func (this *read) ENDPOINT_GetAssignmentGroups(
	input.Bind[requestPresenter.GetAssignmentGroups],
	input.BindInputAuthority,
	input.BindInputTenant,
) endpoint.IEndpoint {

	return this.GET("/{uuid:uuid}/group/list").
		UseMiddleware(
			middleware.Auth(
				hook.AuthRequireTenantAgent,
			),
			// middleware.B>indRequest[requestPresenter.GetAssignmentGroups](
			// 	hook.UseAuthority,
			// 	hook.UseTenantMapping,
			// ),
		).Build()
}
func (this *read) GetAssignmentGroups(
	input *requestPresenter.GetAssignmentGroups,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetAssignmnentGroupsUseCase.Execute(input),
	)
}
