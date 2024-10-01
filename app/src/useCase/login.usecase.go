package usecase

import (
	authServiceAdapter "app/adapter/auth"
	refreshTokenClientPort "app/adapter/refreshTokenClient"
	actionResultService "app/service/actionResult"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"errors"

	"github.com/kataras/iris/v12/mvc"
)

var (
	ERR_NIL_CONTEXT = errors.New("login usecase error: nil context")
)

type (
	ILogIn interface {
		Execute(*requestPresenter.LoginRequest, *responsePresenter.LoginResponse) (mvc.Result, error)
	}

	LogInUseCase struct {
		LogInService       authServiceAdapter.ILogIn
		ActionResult       actionResultService.IActionResult
		RefreshTokenClient refreshTokenClientPort.IRefreshTokenClient
	}
)

func (this *LogInUseCase) Execute(
	input *requestPresenter.LoginRequest,
	output *responsePresenter.LoginResponse,
) (mvc.Result, error) {

	reqContext := input.GetContext()

	if reqContext == nil {

		this.ActionResult.ServeErrorResponse(ERR_NIL_CONTEXT)
	}

	accessToken, refreshToken, err := this.LogInService.Serve(input.Data.Username, input.Data.Password, reqContext)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	err = this.RefreshTokenClient.Write(reqContext, refreshToken)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Data.AccessToken = accessToken
	output.Message = "success"

	this.RefreshTokenClient.Write(reqContext, refreshToken)

	return this.ActionResult.ServeResponse(output)
}
