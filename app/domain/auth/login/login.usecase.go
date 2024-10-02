package loginDomain

import (
	actionResultServicePort "app/port/actionResult"
	authServicePort "app/port/auth"
	refreshTokenClientPort "app/port/refreshTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"
)

var (
	ERR_NIL_CONTEXT = errors.New("login usecase error: nil context")
)

type (
	// ILogIn interface {
	// 	Execute(*requestPresenter.LoginRequest, *responsePresenter.LoginResponse) (mvc.Result, error)
	// }

	LogInUseCase struct {
		usecasePort.UseCase[responsePresenter.LoginResponse]
		LogInService       authServicePort.ILogIn
		ActionResult       actionResultServicePort.IActionResult
		RefreshTokenClient refreshTokenClientPort.IRefreshTokenClient
	}
)

func (this *LogInUseCase) Execute(
	input *requestPresenter.LoginRequest,
) (*responsePresenter.LoginResponse, error) {

	reqContext := input.GetContext()

	if reqContext == nil {

		this.ActionResult.ServeErrorResponse(ERR_NIL_CONTEXT)
	}

	accessToken, refreshToken, err := this.LogInService.Serve(input.Data.Username, input.Data.Password, reqContext)

	if err != nil {

		//return this.ActionResult.ServeErrorResponse(err)
		return nil, err
	}

	err = this.RefreshTokenClient.Write(reqContext, refreshToken)

	if err != nil {

		//return this.ActionResult.ServeErrorResponse(err)
		return nil, err
	}

	output := this.GenerateOutput()

	output.Data.AccessToken = accessToken
	output.Message = "success"

	//this.RefreshTokenClient.Write(reqContext, refreshToken)

	//return this.ActionResult.ServeResponse(output)

	return output, nil
}
