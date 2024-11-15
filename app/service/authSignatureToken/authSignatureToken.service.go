package authSignatureToken

import (
	accessTokenServicePort "app/port/accessToken"
	generalTokenServicePort "app/port/generalToken"
	"app/unitOfWork"

	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	IGeneralToken             = generalTokenServicePort.IGeneralToken
	AuthSignatureTokenService struct {
		unitOfWork.OperationLogger
		AccessTokenManipulator  accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator
	}
)

func (this *AuthSignatureTokenService) Generate(
	TenantUUID uuid.UUID, generalToken IGeneralToken, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	rt, err = this.RefreshTokenManipulator.Generate(TenantUUID, generalToken, ctx)

	if err != nil {

		return
	}

	at, err = this.AccessTokenManipulator.GenerateFor(TenantUUID, generalToken, rt.GetTokenID(), ctx)

	if err != nil {

		return
	}

	at.SetTokenID(rt.GetTokenID())

	return
}

func (this *AuthSignatureTokenService) Rotate(
	oldRefreshToken refreshTokenServicePort.IRefreshToken, oldAccessToken accessTokenServicePort.IAccessToken, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	rt, err = this.RefreshTokenManipulator.Rotate(oldRefreshToken, ctx)

	if err != nil {

		return
	}

	at, err = this.AccessTokenManipulator.GenerateBased(oldAccessToken, ctx)

	if err != nil {

		return
	}

	at.SetTokenID(rt.GetTokenID())

	return
}

func (this *AuthSignatureTokenService) GenerateStrings(
	tenantUUID uuid.UUID, generalToken IGeneralToken, ctx context.Context,
) (at string, rt string, err error) {

	acc, re, err := this.Generate(tenantUUID, generalToken, ctx)

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
	oldRefreshToken refreshTokenServicePort.IRefreshToken, oldAccessToken accessTokenServicePort.IAccessToken, ctx context.Context,
) (at string, rt string, err error) {

	acc, re, err := this.Rotate(oldRefreshToken, oldAccessToken, ctx)

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
