package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AssignmentController struct {
		CreateAssignmentUseCase      usecase.ICreateAssignment
		GetSingleAssignmentUseCase   usecase.IGetSingleAssignment
		CreateAssignmentGroupUseCase usecase.ICreateAssignmentGroup
	}
)

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
