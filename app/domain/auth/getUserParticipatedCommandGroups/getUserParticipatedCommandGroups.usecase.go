package getUserParticipatedCommandGroupsDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GetUserParticipatedCommandGroupsUseCase struct {
		usecasePort.UseCase[requestPresenter.GetUserParticipatedCommandGroups, responsePresenter.GetUserParticipatedCommandGroups]
		GetUserParticipatedCommandGroupService authServicePort.IGetUserParticipatedCommandGroups
	}
)

func (this *GetUserParticipatedCommandGroupsUseCase) Execute(
	input *requestPresenter.GetUserParticipatedCommandGroups,
) (*responsePresenter.GetUserParticipatedCommandGroups, error) {

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	data, err := this.GetUserParticipatedCommandGroupService.Serve(
		input.GetTenantUUID(), *input.UserUUID, input.GetContext(),
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
