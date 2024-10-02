package authSignatureToken

import (
	accessTokenServicePort "app/port/accessToken"
	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	AuthSignatureTokenService struct {
		AccessTokenManipulator  accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator
	}
)

func (this *AuthSignatureTokenService) Generate(
	userUUID uuid.UUID, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	rt, err = this.RefreshTokenManipulator.Generate(userUUID, ctx)

	if err != nil {

		return
	}

	at, err = this.AccessTokenManipulator.GenerateByUserUUID(userUUID, rt.GetTokenID(), ctx)

	if err != nil {

		return
	}

	return
}

func (this *AuthSignatureTokenService) Rotate(
	refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	rt, err = this.RefreshTokenManipulator.Rotate(refreshToken, ctx)

	if err != nil {

		return
	}

	at, err = this.AccessTokenManipulator.GenerateByUserUUID(refreshToken.GetUserUUID(), rt.GetTokenID(), ctx)

	if err != nil {

		return
	}

	return
}

func (this *AuthSignatureTokenService) GenerateStrings(
	userUUID uuid.UUID, ctx context.Context,
) (at string, rt string, err error) {

	acc, re, err := this.Generate(userUUID, ctx)

	if err != nil {

		return
	}

	at, err = this.AccessTokenManipulator.SignString(acc)

	if err != nil {

		return
	}

	rt, err = this.RefreshTokenManipulator.SignString(re)

	if err != nil {

		return
	}

	return
}

func (this *AuthSignatureTokenService) RotateStrings(
	refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context,
) (at string, rt string, err error) {

	acc, re, err := this.Rotate(refreshToken, ctx)

	if err != nil {

		return
	}

	at, err = this.AccessTokenManipulator.SignString(acc)

	if err != nil {

		return
	}

	rt, err = this.RefreshTokenManipulator.SignString(re)

	if err != nil {

		return
	}

	return
}
