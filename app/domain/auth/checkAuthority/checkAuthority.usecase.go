package checkAuthorityDomain

import (
	"app/internal/common"
	accessTokenServicePort "app/port/accessToken"
	accessTokenClientPort "app/port/accessTokenClient"
	authServicePort "app/port/auth"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	usecasePort "app/port/usecase"
	"context"
	"errors"
	"fmt"
)

const (
	USE_CASE_NAME = "CheckAuthorityUseCase"
)

type (
	CheckAuthorityUseCase struct {
		usecasePort.UserSessionCacheUseCase
		CheckAuthorityService  authServicePort.ICheckAuthority
		AccessTokenClient      accessTokenClientPort.IAccessTokenClient
		AccessTokenManipulator accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenIDProvider refreshTokenIdServicePort.IRefreshTokenIDProvider
	}
)

func (this *CheckAuthorityUseCase) Execute(
	ctx context.Context,
) error {

	accessToken, err := this.AccessTokenClient.Read(ctx)

	switch {
	case err != nil:
		return err
	case accessToken == nil:
		return errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("%s error: no access token", USE_CASE_NAME))
	case accessToken.Expired():
		return errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("%s error: access token expires", USE_CASE_NAME))
	}

	// access token and refresh token share the same token id
	inBlackList, err := this.RefreshTokenBlackList.Has(accessToken.GetTokenID(), ctx)

	if err != nil {

		return err
	}

	if inBlackList {

		return errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("%s error: access token deactivated", USE_CASE_NAME))
	}

	generalTokenID, _, err := this.RefreshTokenIDProvider.Extract(accessToken.GetTokenID())

	if err != nil {

		return errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("%s error:", USE_CASE_NAME), err)
	}

	err = this.CheckAuthorityService.Serve(
		accessToken.GetTenantUUID(), accessToken.GetUserUUID(), generalTokenID, ctx,
	)

	if err != nil {

		return err
	}

	return nil
}
