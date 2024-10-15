package getAssignmentsDomain

import (
	"app/internal/common"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GetAssignmentUseCase struct {
		usecasePort.UseCase[requestPresenter.GetAssignments, responsePresenter.GetAssignments]
		GetAssignmentService assignmentServicePort.IGetAssignments
	}
)

func (this *GetAssignmentUseCase) Execute(
	input *requestPresenter.GetAssignments,
) (*responsePresenter.GetAssignments, error) {

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	if !input.GetAuthority().IsTenantAgent() {

		return nil, this.ErrorWithContext(
			input, common.ERR_FORBIDEN,
		)
	}

	data, err := this.GetAssignmentService.Serve(
		input.GetTenantUUID(),
		input,
		input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	output := this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}
