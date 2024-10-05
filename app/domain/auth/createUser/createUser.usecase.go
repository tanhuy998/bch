package createUserDomain

import (
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"fmt"
)

type (
	// ICreateUser interface {
	// 	Execute(
	// 		input *requestPresenter.CreateUserRequestPresenter,
	// 		output *responsePresenter.CreateUserPresenter,
	// 	) (mvc.Result, error)
	// }

	CreateUserUsecase struct {
		usecasePort.UseCase[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter]
		CreateUserService authServicePort.ICreateUser
	}
)

func (this *CreateUserUsecase) Execute(
	input *requestPresenter.CreateUserRequestPresenter,
) (*responsePresenter.CreateUserPresenter, error) {

	//_, err := this.CreateUserService.Serve(input.Data.Username, input.Data.Password, input.Data.Name, input.GetContext())

	newUser := &model.User{
		Username:   input.Data.Username,
		PassWord:   input.Data.Password,
		Name:       input.Data.Name,
		TenantUUID: libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID()),
		CreatedBy:  libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID()),
	}

	data, err := this.CreateUserService.CreateByModel(newUser, input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	if data == nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(fmt.Errorf("cannot create user, try again")),
		)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = data

	return output, nil
}
