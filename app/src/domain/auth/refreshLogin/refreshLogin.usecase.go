package usecase

import (
	authServiceAdapter "app/adapter/auth"
	refreshTokenServicePort "app/adapter/refreshToken"
	refreshTokenClientPort "app/adapter/refreshTokenClient"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kataras/iris/v12/mvc"
)

var (
	ERR_REFRESH_NO_CONTEXT = errors.New("refresh login usecase: no context")
)

type (
	IRefreshLogin interface {
		Execute(
			input *requestPresenter.RefreshLoginRequest,
			output *responsePresenter.RefreshLoginResponse,
		) (mvc.Result, error)
	}

	RefreshLoginUseCase struct {
		RefreshLoginService authServiceAdapter.IRefreshLogin
		ActionResult        actionResultService.IActionResult
		RefreshTokenClient  refreshTokenClientPort.IRefreshTokenClient
	}
)

func (this *RefreshLoginUseCase) Execute(
	input *requestPresenter.RefreshLoginRequest,
	output *responsePresenter.RefreshLoginResponse,
) (mvc.Result, error) {

	reqCtx := input.GetContext()

	if reqCtx == nil {

		return this.ActionResult.ServeErrorResponse(ERR_REFRESH_NO_CONTEXT)
	}

	refreshToken_str := this.RefreshTokenClient.Read(reqCtx)

	newAccessTokenString, newRefreshTokenString, err := this.RefreshLoginService.Serve(input.Data.AccessToken, refreshToken_str, reqCtx)

	switch err {
	case nil:
		output.Message = "success"
		output.Data = &responsePresenter.RefreshLoginData{
			newAccessTokenString,
		}
	case authServiceAdapter.ERR_REFRESH_TOKEN_EXPIRE, refreshTokenServicePort.ERR_REFRESH_TOKEN_BLACK_LIST:
		output.Message = err.Error()
		raw_content, _ := json.Marshal(output)
		return this.ActionResult.Prepare().SetCode(http.StatusForbidden).SetContent(raw_content).Done()
	default:
		return this.ActionResult.ServeErrorResponse(err)
	}

	err = this.RefreshTokenClient.Write(reqCtx, newRefreshTokenString)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	return this.ActionResult.ServeResponse(output)
}
