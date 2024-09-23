package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AssignmentController struct {
		CreateAssignmentUseCase usecase.ICreateAssignment
	}
)

func (this *AssignmentController) CreateAssignment(
	input *requestPresenter.CreateAssigmentRequest,
	output *responsePresenter.CreateAssignmentResponse,
) (mvc.Result, error) {

	return this.CreateAssignmentUseCase.Execute(input, output)
}
