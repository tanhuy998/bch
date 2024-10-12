package modifyAssignmentDomain

import (
	"app/internal/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	ModifyAssignmentUseCase struct {
		usecasePort.UseCase[requestPresenter.ModifyAssignment, responsePresenter.ModifyAssignment]
		ModifyAssignmentService assignmentServicePort.IModifyAssignment
	}
)

func (this *ModifyAssignmentUseCase) Execute(
	input *requestPresenter.ModifyAssignment,
) (*responsePresenter.ModifyAssignment, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	dataModel := &model.Assignment{
		Title:      input.Data.Title,
		Deadline:   input.Data.Deadline,
		Desciption: input.Data.Description,
	}

	err := this.ModifyAssignmentService.Serve(
		input.GetTenantUUID(), *input.AssignmentUUID, dataModel, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"

	return output, nil
}
