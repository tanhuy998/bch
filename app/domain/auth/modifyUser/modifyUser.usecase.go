package modifyUserDomain

import (
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	// IModifyUser interface {
	// 	Execute(
	// 		input *requestPresenter.ModifyUserRequest,
	// 		output *responsePresenter.ModifyUserResponse,
	// 	) (mvc.Result, error)
	// }

	ModifyUserUseCase struct {
		usecasePort.UseCase[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse]
		ModifyUser authServicePort.IModifyUser
	}
)

func (this *ModifyUserUseCase) Execute(
	input *requestPresenter.ModifyUserRequest,
) (*responsePresenter.ModifyUserResponse, error) {

	dataModel := &model.User{
		Name:     input.Data.Name,
		PassWord: input.Data.Password,
	}

	err := this.ModifyUser.Serve(*input.UserUUID, dataModel, input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()
	output.Message = "success"

	return output, nil
}
