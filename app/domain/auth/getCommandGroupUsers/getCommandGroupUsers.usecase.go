package getCommandGroupUsersDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GetCommandGroupUsersUseCase struct {
		usecasePort.UseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse]
		GetCommandGroupUserService authServicePort.IGetCommandGroupUsers
	}
)

func (this *GetCommandGroupUsersUseCase) Execute(
	input *requestPresenter.GetGroupUsersRequest,
) (*responsePresenter.GetGroupUsersResponse, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	data, err := this.GetCommandGroupUserService.Serve(
		input.GetTenantUUID(), *input.GroupUUID, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}
