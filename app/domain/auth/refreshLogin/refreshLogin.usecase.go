package refreshLoginDomain

import (
	libError "app/internal/lib/error"
	authServicePort "app/port/auth"
	refreshTokenClientPort "app/port/refreshTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"
	"fmt"

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
		usecasePort.UseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
		RefreshLoginService authServicePort.IRefreshLogin
		RefreshTokenClient  refreshTokenClientPort.IRefreshTokenClient
	}
)

func (this *RefreshLoginUseCase) Execute(
	input *requestPresenter.RefreshLoginRequest,
) (*responsePresenter.RefreshLoginResponse, error) {

	reqCtx := input.GetContext()

	if reqCtx == nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(fmt.Errorf("refreshLoginUseCase: nil context given")),
		)
	}

	refreshToken_str := this.RefreshTokenClient.Read(reqCtx)

	newAccessTokenString, newRefreshTokenString, err := this.RefreshLoginService.Serve(input.Data.AccessToken, refreshToken_str, reqCtx)
	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	err = this.RefreshTokenClient.Write(reqCtx, newRefreshTokenString)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = &responsePresenter.RefreshLoginData{
		newAccessTokenString,
	}

	return output, nil
}
