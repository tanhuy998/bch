package refreshLoginDomain

import (
	"app/internal/common"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
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
		RefreshLoginService    authServicePort.IRefreshLogin
		AccessTokenManipulator accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenClient     refreshTokenClientPort.IRefreshTokenClient
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

	refreshToken, err := this.RefreshTokenClient.Read(reqCtx)

	if err != nil {

		return nil, err
	}

	if refreshToken == nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	accessToken, err := this.AccessTokenManipulator.Read(input.Data.AccessToken)

	if err != nil {

		return nil, err
	}

	newAccessToken, newRefreshToken, err := this.RefreshLoginService.Serve(accessToken, refreshToken, reqCtx)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	err = this.RefreshTokenClient.Write(reqCtx, newRefreshToken)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	at, err := this.AccessTokenManipulator.SignString(newAccessToken)

	if err != nil {

		return nil, err
	}

	output.Message = "success"
	output.Data = &responsePresenter.RefreshLoginData{
		at,
	}

	return output, nil
}
