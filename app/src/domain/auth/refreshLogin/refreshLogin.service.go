package refreshLogin

import (
	"app/src/internal/common"
	accessTokenServicePort "app/src/port/accessToken"
	authSignatureTokenPort "app/src/port/authSignatureToken"
	refreshTokenServicePort "app/src/port/refreshToken"
	refreshTokenBlackListServicePort "app/src/port/refreshTokenBlackList"
	"context"
	"errors"
)

var (
	ERR_BAD_TOKEN_PAIR = errors.New("refresh login error: bad token pair")
)

const (
	SERVICE_NAME = "(RefreshLoginService)"
)

type (
	RefreshLoginService struct {
		RefreshTokenBlackList      refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator
		RefreshTokenManipulator    refreshTokenServicePort.IRefreshTokenManipulator
		AccessTokenManipulator     accessTokenServicePort.IAccessTokenManipulator
		AuthSignatureTokenProvider authSignatureTokenPort.IAuthSignatureProvider
	}
)

func (this *RefreshLoginService) Serve(
	inputAT string, inputRT string, reqCtx context.Context,
) (at string, rt string, err error) {

	accessToken, err := this.AccessTokenManipulator.Read(inputAT)

	if err != nil {

		return
	}

	if !accessToken.Expired() {

		err = errors.New("(RefreshLoginService) access token not expired") // authServiceAdapter.ERR_ACCESS_TOKEN_NOT_EXPIRE
		return
	}

	refreshToken, err := this.RefreshTokenManipulator.Read(inputRT)

	switch {
	case err != nil:
		return
	case refreshToken.Expired():
		err = errors.New("(RefreshLoginService) refresh Token expired") // authServiceAdapter.ERR_REFRESH_TOKEN_EXPIRE
		return
	case refreshToken.GetTokenID() != accessToken.GetTokenID(), refreshToken.GetUserUUID() != accessToken.GetUserUUID():
		err = errors.Join(common.ERR_UNAUTHORIZED, errors.New("(RefreshLoginService) signatures mismatch")) // ERR_BAD_TOKEN_PAIR
		return
	}

	refreshTokenRevoked, err := this.RefreshTokenBlackList.Has(refreshToken.GetTokenID(), reqCtx)

	switch {
	case err != nil:
		return
	case refreshTokenRevoked:
		err = errors.Join(common.ERR_FORBIDEN, errors.New("(RefreshLoginService) refresh token revoked")) //authServiceAdapter.ERR_REFRESH_TOKEN_EXPIRE
		return
		// case refreshToken.GetUserUUID() != accessToken.GetUserUUID():
		// 	err = errors.New() //authServiceAdapter.ERR_REFESH_LOGIN_INVALID_ACCESS_TOKEN
		// 	return
	}

	return this.AuthSignatureTokenProvider.RotateStrings(refreshToken, reqCtx)
}
