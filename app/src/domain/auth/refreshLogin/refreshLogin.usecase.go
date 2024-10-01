package refreshLogin

import (
	"app/src/internal/common"
	actionResultServicePort "app/src/port/actionResult"
	authServicePort "app/src/port/auth"
	refreshTokenClientPort "app/src/port/refreshTokenClient"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
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
		RefreshLoginService authServicePort.IRefreshLogin
		ActionResult        actionResultServicePort.IActionResult
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

	switch {
	case errors.Is(err, common.ERR_FORBIDEN):
		output.Message = err.Error()
		raw_content, _ := json.Marshal(output)
		return this.ActionResult.Prepare().SetCode(http.StatusForbidden).SetContent(raw_content).Done()
	case errors.Is(err, common.ERR_UNAUTHORIZED):
		output.Message = err.Error()
		raw_content, _ := json.Marshal(output)
		return this.ActionResult.Prepare().SetCode(http.StatusUnauthorized).SetContent(raw_content).Done()
	case err != nil:
		return this.ActionResult.ServeErrorResponse(err)
	}

	err = this.RefreshTokenClient.Write(reqCtx, newRefreshTokenString)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"
	output.Data = &responsePresenter.RefreshLoginData{
		newAccessTokenString,
	}

	return this.ActionResult.ServeResponse(output)
}
