package createAssignmentDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	CreateAssignmentUseCase struct {
		usecasePort.UseCase[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse]
		CreateAssignmentService assignmentServicePort.ICreateAssignment
	}
)

func (this *CreateAssignmentUseCase) Execute(
	input *requestPresenter.CreateAssigmentRequest,
) (*responsePresenter.CreateAssignmentResponse, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	inputData := input.Data

	model := &model.Assignment{
		Title:      inputData.Title,
		CreatedBy:  libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID()),
		Desciption: inputData.Desciption,
		Deadline:   inputData.DeadLine,
	}

	data, err := this.CreateAssignmentService.Serve(
		input.GetTenantUUID(), model, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = data

	return output, nil
}
