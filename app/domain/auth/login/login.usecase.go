package loginDomain

import (
	libError "app/internal/lib/error"
	authServicePort "app/port/auth"
	refreshTokenClientPort "app/port/refreshTokenClient"
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
	// ILogIn interface {
	// 	Execute(*requestPresenter.LoginRequest, *responsePresenter.LoginResponse) (mvc.Result, error)
	// }

	LogInUseCase struct {
		usecasePort.UseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse]
		LogInService       authServicePort.ILogIn
		RefreshTokenClient refreshTokenClientPort.IRefreshTokenClient
	}
)

func (this *LogInUseCase) Execute(
	input *requestPresenter.LoginRequest,
) (*responsePresenter.LoginResponse, error) {

	reqContext := input.GetContext()

	if reqContext == nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(fmt.Errorf("loginUseCase: nil context given to usecase")),
		)
	}

	accessToken, refreshToken, err := this.LogInService.Serve(input.Data.Username, input.Data.Password, reqContext)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	err = this.RefreshTokenClient.Write(reqContext, refreshToken)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Data.AccessToken = accessToken
	output.Message = "success"
	return output, nil
}
