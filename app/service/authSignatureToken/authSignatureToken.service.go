package authSignatureToken

import (
	accessTokenServicePort "app/port/accessToken"
	generalTokenServicePort "app/port/generalToken"
	"fmt"

	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	IGeneralToken             = generalTokenServicePort.IGeneralToken
	AuthSignatureTokenService struct {
		AccessTokenManipulator  accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator
	}
)

func (this *AuthSignatureTokenService) Generate(
	TenantUUID uuid.UUID, generalToken IGeneralToken, ctx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	rt, err = this.RefreshTokenManipulator.Generate(TenantUUID, generalToken, ctx)
	fmt.Println(5)
	if err != nil {

		return
	}

	at, err = this.AccessTokenManipulator.GenerateFor(TenantUUID, generalToken, rt.GetTokenID(), ctx)
	fmt.Println(6, err)
	if err != nil {

		return
	}
	fmt.Println(7)
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
