package logoutDomain

import (
	"app/internal/common"
	libError "app/internal/lib/error"
	accessTokenClientPort "app/port/accessTokenClient"
	authServicePort "app/port/auth"
	refreshTokenServicePort "app/port/refreshToken"
	refreshTokenClientPort "app/port/refreshTokenClient"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type (
	LogoutUseCase struct {
		usecasePort.MongoUserSessionCacheUseCase[responsePresenter.Logout]
		usecasePort.UseCase[requestPresenter.Logout, responsePresenter.Logout]
		AccessTokenClientService  accessTokenClientPort.IAccessTokenClient
		RefreshTokenClientService refreshTokenClientPort.IRefreshTokenClient
		RefreshTokenIDProvider    refreshTokenIdServicePort.IRefreshTokenIDProvider
		LogoutService             authServicePort.ILogout
	}
)

func (this *LogoutUseCase) Execute(
	input *requestPresenter.Logout,
) (output *responsePresenter.Logout, err error) {

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("nil access token given")),
		)
	}

	accessToken, err := this.AccessTokenClientService.Read(input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	refreshToken, err := this.RefreshTokenClientService.Read(input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	output, err = this.ModifyUserSession(
		input.GetContext(),
		func(ctx context.Context) (ret *responsePresenter.Logout, err error) {

			defer func() {

				if err != nil {

					return
				}

				err = this.markRefreshToken(refreshToken, input.GetContext())

				if err != nil {

					output = nil
					return
				}

				err = this.RemoveUserSession(
					ctx, refreshToken.GetUserUUID(),
				)

				if err != nil {

					output = nil
					return
				}
			}()

			err = this.LogoutService.Serve(refreshToken, accessToken, ctx)

			if err != nil {

				return nil, this.ErrorWithContext(
					input, err,
				)
			}

			output := this.GenerateOutput()
			output.Message = "success"
			output.SetHTTPStatus(http.StatusAccepted)

			return output, nil
		},
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	return output, nil
}

func (this *LogoutUseCase) markRefreshToken(refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context) error {

	expireTime := refreshToken.GetExpireTime()

	var (
		setted bool
		err    error
	)

	if expireTime == nil {

		setted, err = this.RefreshTokenBlackList.Set(refreshToken.GetTokenID(), struct{}{}, ctx)

	} else {

		err = this.RefreshTokenBlackList.SetWithExpire(refreshToken.GetTokenID(), struct{}{}, *expireTime, ctx)
	}

	if err != nil {

		return err
	}

	if !setted {

		return libError.NewInternal(fmt.Errorf("cannot set refresh token id to black list"))
	}

	err = this.RefreshTokenClientService.Remove(ctx)

	if err != nil {

		return err
	}

	return nil
}
