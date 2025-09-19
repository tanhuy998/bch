package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
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
		crud.ReadEndpointCurator
		common.Controller
		GetSingleAssignmentUseCase  usecasePort.IUseCase[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse]
		GetAssignmentsUseCase       usecasePort.IUseCase[requestPresenter.GetAssignments, responsePresenter.GetAssignments[model.Assignment]]
		GetAssignmnentGroupsUseCase usecasePort.IUseCase[requestPresenter.GetAssignmentGroups, responsePresenter.GetAssignmentGroups[model.AssignmentGroup]]
	}
)

func (this *read) ENDPOINT_GetAssignments(
	input.Bind[requestPresenter.GetAssignments],

	//auth.Authorize[constraint.TenantAgent],
) endpoint.IEndpoint {

	return this.GET("/").
		BuildAction(
			func(input *requestPresenter.GetAssignments) (mvc.Result, error) {

				return this.ResultOf(
					this.GetAssignmentsUseCase.Execute(input),
				)
			},
		)
}

func (this *read) ENDPOINT_GetSingleAssignment(
	input.Bind[requestPresenter.GetSingleAssignmentRequest],
) endpoint.IEndpoint {

	return this.GET("/{uuid:uuid}").
		BuildAction(
			func(input *requestPresenter.GetSingleAssignmentRequest) (mvc.Result, error) {

				return this.ResultOf(
					this.GetSingleAssignmentUseCase.Execute(input),
				)
			},
		)
}

func (this *read) ENDPOINT_GetAssignmentGroups(
	input.Bind[requestPresenter.GetAssignmentGroups],
	auth.Authorize[constraint.TenantAgent],
) endpoint.IEndpoint {
	return this.GET("/{uuid:uuid}/group/list").
		BuildAction(
			func(input *requestPresenter.GetAssignmentGroups) (mvc.Result, error) {

				return this.ResultOf(
					this.GetAssignmnentGroupsUseCase.Execute(input),
				)
			},
		)
}
