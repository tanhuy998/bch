package generateAuthSignatureTokensDomain

import (
	accessTokenServicePort "app/port/accessToken"
	generalTokenServicePort "app/port/generalToken"
	refreshTokenServicePort "app/port/refreshToken"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	IGeneralToken             = generalTokenServicePort.IGeneralToken
	AuthSignatureTokenService struct {
		AccessTokenManipulator  accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator
	}
)

func (this *AuthSignatureTokenService) Serve(
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

func (this *AuthSignatureTokenService) ServeStrings(
	tenantUUID uuid.UUID, generalToken IGeneralToken, ctx context.Context,
) (at string, rt string, err error) {

	acc, re, err := this.Serve(tenantUUID, generalToken, ctx)

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
