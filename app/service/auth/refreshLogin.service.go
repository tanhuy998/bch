package authService

import (
	accessTokenServicePort "app/adapter/accessToken"
	authServiceAdapter "app/adapter/auth"
	authSignatureTokenPort "app/adapter/authSignatureToken"
	refreshTokenServicePort "app/adapter/refreshToken"
	refreshTokenBlackListServicePort "app/adapter/refreshTokenBlackList"
	"context"
	"errors"
)

var (
	ERR_BAD_TOKEN_PAIR = errors.New("refresh login error: bad token pair")
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

		err = authServiceAdapter.ERR_ACCESS_TOKEN_NOT_EXPIRE
		return
	}

	refreshToken, err := this.RefreshTokenManipulator.Read(inputRT)

	switch {
	case err != nil:
		return
	case refreshToken.Expired():
		err = authServiceAdapter.ERR_REFRESH_TOKEN_EXPIRE
		return
	}

	refreshTokenRevoked, err := this.RefreshTokenBlackList.Has(refreshToken.GetTokenID(), reqCtx)

	switch {
	case err != nil:
		return
	case refreshTokenRevoked:
		err = authServiceAdapter.ERR_REFRESH_TOKEN_EXPIRE
		return
	case refreshToken.GetTokenID() != accessToken.GetTokenID():
		err = ERR_BAD_TOKEN_PAIR
		return
	case refreshToken.GetUserUUID() != accessToken.GetUserUUID():
		err = authServiceAdapter.ERR_REFESH_LOGIN_INVALID_ACCESS_TOKEN
		return
	}

	return this.AuthSignatureTokenProvider.RotateStrings(refreshToken, reqCtx)
}
