package getUserParticipatedCommandGroupsDomain

import (
	"app/internal/common"
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GetUserParticipatedCommandGroupsUseCase struct {
		usecasePort.UseCase[requestPresenter.GetUserParticipatedCommandGroups, responsePresenter.GetUserParticipatedCommandGroups[model.CommandGroup]]
		GetUserParticipatedCommandGroupService authServicePort.IGetUserParticipatedCommandGroups[model.CommandGroup]
	}
)

func (this *GetUserParticipatedCommandGroupsUseCase) Execute(
	input *requestPresenter.GetUserParticipatedCommandGroups,
) (*responsePresenter.GetUserParticipatedCommandGroups[model.CommandGroup], error) {

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	// data, err := this.GetUserParticipatedCommandGroupService.Serve(
	// 	input.GetTenantUUID(), *input.UserUUID, input.GetContext(),
	// )

	data, err := this.GetUserParticipatedCommandGroupService.Serve(
		input,
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
