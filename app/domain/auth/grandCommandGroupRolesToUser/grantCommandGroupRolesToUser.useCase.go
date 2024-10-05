package grantCommandGroupRoleToUserDomain

import (
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	// IGrantCommandGroupRolesToUser interface {
	// 	Execute(
	// 		input *requestPresenter.GrantCommandGroupRolesToUserRequest,
	// 		output *responsePresenter.GrantCommandGroupRolesToUserResponse,
	// 	) (mvc.Result, error)
	// }

	GrantCommandGroupRolesToUserUseCase struct {
		usecasePort.UseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse]
		GrantCommandGroupRolesToUserService authServicePort.IGrantCommandGroupRolesToUser
	}
)

func (this *GrantCommandGroupRolesToUserUseCase) Execute(
	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
) (*responsePresenter.GrantCommandGroupRolesToUserResponse, error) {

	err := this.GrantCommandGroupRolesToUserService.Serve(*input.GroupUUID, *input.UserUUID, input.Data, input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()
	output.Message = "success"

	return output, nil
}
