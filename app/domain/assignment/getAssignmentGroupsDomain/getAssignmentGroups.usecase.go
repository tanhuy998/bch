package getAssignmentGroupsDomain

import (
	"app/internal/common"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GetAssignmentGroupsUseCase struct {
		usecasePort.UseCase[requestPresenter.GetAssignmentGroups, responsePresenter.GetAssignmentGroups]
		GetAssignmentGroupService assignmentServicePort.IGetAssignmentGroups
	}
)

func (this *GetAssignmentGroupsUseCase) Execute(
	input *requestPresenter.GetAssignmentGroups,
) (*responsePresenter.GetAssignmentGroups, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	if !input.IsTenantAgent() {

		return nil, common.ERR_FORBIDEN
	}

	data, err := this.GetAssignmentGroupService.Serve(
		input.GetTenantUUID(), *input.AssignmentUUID, input.GetContext(),
	)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = data

	return output, nil
}
