package authenticateCredentialsDomain

import (
	libError "app/internal/lib/error"
	authGenServicePort "app/port/authGenService"
	generalTokenClientServicePort "app/port/generalTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"
	"fmt"
)

var (
	ERR_NIL_CONTEXT = errors.New("login usecase error: nil context")
)

type (
	AuthenticateCredentialsUseCase struct {
		usecasePort.UseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse]
		LogInService              authGenServicePort.IAuthenticateCrdentials
		GeneralTokenClientService generalTokenClientServicePort.IGeneralTokenClient
	}
)

func (this *AuthenticateCredentialsUseCase) Execute(
	input *requestPresenter.LoginRequest,
) (*responsePresenter.LoginResponse, error) {

	reqContext := input.GetContext()

	if reqContext == nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(fmt.Errorf("AuthenticateCredentialsUseCase: nil context given to usecase")),
		)
	}

	generalToken, err := this.LogInService.Serve(input.Data.Username, input.Data.Password, reqContext)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	err = this.GeneralTokenClientService.Write(reqContext, generalToken)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	return output, nil
}
