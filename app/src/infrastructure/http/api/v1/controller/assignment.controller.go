package controller

import (
	"app/src/infrastructure/http/common"
	"app/src/infrastructure/http/middleware"
	"app/src/infrastructure/http/middleware/middlewareHelper"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	usecase "app/src/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AssignmentController struct {
		*common.Controller
		CreateAssignmentUseCase      usecase.ICreateAssignment
		GetSingleAssignmentUseCase   usecase.IGetSingleAssignment
		CreateAssignmentGroupUseCase usecase.ICreateAssignmentGroup
	}
)

func (this *AssignmentController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	activator.Handle(
		"GET", "/{uuid:uuid}", "GetSingleAssignment",
		middleware.BindPresenters[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"POST", "/", "CreateAssignment",
		middleware.BindPresenters[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"POST", "/{assignmentUUID:uuid}/group", "CreateAssignmentGroup",
		middleware.BindPresenters[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)
}

func (this *AssignmentController) GetSingleAssignment(
	input *requestPresenter.GetSingleAssignmentRequest,
	output *responsePresenter.GetSingleAssignmentResponse,
) (mvc.Result, error) {

	return this.GetSingleAssignmentUseCase.Execute(input, output)
}

func (this *AssignmentController) CreateAssignment(
	input *requestPresenter.CreateAssigmentRequest,
	output *responsePresenter.CreateAssignmentResponse,
) (mvc.Result, error) {

	return this.CreateAssignmentUseCase.Execute(input, output)
}

func (this *AssignmentController) CreateAssignmentGroup(
	input *requestPresenter.CreateAssignmentGroupRequest,
	output *responsePresenter.CreateAssignmentGroupResponse,
) (mvc.Result, error) {

	return this.CreateAssignmentGroupUseCase.Execute(input, output)
}
