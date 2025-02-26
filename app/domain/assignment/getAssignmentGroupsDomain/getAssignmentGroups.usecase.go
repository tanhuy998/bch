package getAssignmentGroupsDomain

import (
	"app/domain"
	"app/internal/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GetAssignmentGroupsUseCase struct {
		usecasePort.UseCase[requestPresenter.GetAssignmentGroups, responsePresenter.GetAssignmentGroups[model.AssignmentGroup]]
		GetAssignmentGroupService assignmentServicePort.IGetAssignmentGroups[domain.PaginateCursorType, model.AssignmentGroup]
	}
)

func (this *GetAssignmentGroupsUseCase) Execute(
	input *requestPresenter.GetAssignmentGroups,
) (ret *responsePresenter.GetAssignmentGroups[model.AssignmentGroup], err error) {

	// defer func() {

	// 	this.WrapResults(&input, &ret, &err)
	// }()

	switch {
	case !input.IsValidTenantUUID():
		return nil, common.ERR_UNAUTHORIZED
	case !input.IsTenantAgent():
		return nil, common.ERR_FORBIDEN
	}

	// if !input.IsValidTenantUUID() {

	// 	return nil, common.ERR_UNAUTHORIZED
	// }

	// if !input.IsTenantAgent() {

	// 	return nil, common.ERR_FORBIDEN
	// }

	data, err := this.GetAssignmentGroupService.Serve(
		input.GetTenantUUID(), *input.AssignmentUUID, input, input.GetContext(),
	)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = data

	return output, nil
}
