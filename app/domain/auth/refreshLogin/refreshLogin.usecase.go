package refreshLoginDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	refreshTokenClientPort "app/port/refreshTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"

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
		usecasePort.UseCase[responsePresenter.RefreshLoginResponse]
		RefreshLoginService authServicePort.IRefreshLogin
		RefreshTokenClient  refreshTokenClientPort.IRefreshTokenClient
	}
)

func (this *RefreshLoginUseCase) Execute(
	input *requestPresenter.RefreshLoginRequest,
) (*responsePresenter.RefreshLoginResponse, error) {

	reqCtx := input.GetContext()

	if reqCtx == nil {

		//return this.ActionResult.ServeErrorResponse(ERR_REFRESH_NO_CONTEXT)

		return nil, errors.Join(common.ERR_INTERNAL, errors.New("no context in presenter"))
	}

	refreshToken_str := this.RefreshTokenClient.Read(reqCtx)

	newAccessTokenString, newRefreshTokenString, err := this.RefreshLoginService.Serve(input.Data.AccessToken, refreshToken_str, reqCtx)

	// switch {
	// case errors.Is(err, common.ERR_FORBIDEN):
	// 	output.Message = err.Error()
	// 	raw_content, _ := json.Marshal(output)
	// 	return this.ActionResult.Prepare().SetCode(http.StatusForbidden).SetContent(raw_content).Done()
	// case errors.Is(err, common.ERR_UNAUTHORIZED):
	// 	output.Message = err.Error()
	// 	raw_content, _ := json.Marshal(output)
	// 	return this.ActionResult.Prepare().SetCode(http.StatusUnauthorized).SetContent(raw_content).Done()
	// case err != nil:
	// 	return this.ActionResult.ServeErrorResponse(err)
	// }

	if err != nil {

		return nil, err
	}

	err = this.RefreshTokenClient.Write(reqCtx, newRefreshTokenString)

	if err != nil {

		//return this.ActionResult.ServeErrorResponse(err)

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = &responsePresenter.RefreshLoginData{
		newAccessTokenString,
	}

	// return this.ActionResult.ServeResponse(output)

	return output, nil
}
