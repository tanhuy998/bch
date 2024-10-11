package grantCommandGroupRoleToUserDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GrantCommandGroupRolesToUserUseCase struct {
		usecasePort.UseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse]
		GrantCommandGroupRolesToUserService authServicePort.IGrantCommandGroupRolesToUser
	}
)

func (this *GrantCommandGroupRolesToUserUseCase) Execute(
	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
) (*responsePresenter.GrantCommandGroupRolesToUserResponse, error) {

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	auth := input.GetAuthority()

	err := this.GrantCommandGroupRolesToUserService.Serve(
		input.GetTenantUUID(), *input.GroupUUID, *input.UserUUID, input.Data, auth.GetUserUUID(), input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()
	output.Message = "success"

	return output, nil
}
