package authenticatorService

import (
	accessTokenServicePort "app/port/accessToken"
	accessTokenClientPort "app/port/accessTokenClient"
	"app/port/authenticatorServicePort"
	usecasePort "app/port/usecase"
	"context"
)

type (
	AuthenticatorService struct {
		AccessTokenClient            accessTokenClientPort.IAccessTokenClient
		CheckAuthoritySessionUseCase usecasePort.IMiddlewareUseCase
	}
)

func (this *AuthenticatorService) Authorize(ctx context.Context, policies ...authenticatorServicePort.IAuthorizePolicy) error {

	accessToken, err := this.VerifyAccessToken(ctx)

	if err != nil {

		return err
	}

	for _, p := range policies {

		err = p.Apply(accessToken)

		if err != nil {

			return err
		}
	}

	return nil
}

func (this *AuthenticatorService) VerifyAccessToken(ctx context.Context) (accessToken accessTokenServicePort.IAccessToken, err error) {

	err = this.CheckAuthoritySessionUseCase.Execute(ctx)

	if err != nil {

		return
	}

	return this.AccessTokenClient.Read(ctx)
}
