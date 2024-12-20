package rotateSignaturesDomain

import (
	"app/internal/common"
	accessTokenServicePort "app/port/accessToken"
	authSignatureTokenPort "app/port/authSignatureToken"
	refreshTokenServicePort "app/port/refreshToken"
	"app/unitOfWork"
	"context"
	"errors"
	"fmt"
)

var (
	ERR_BAD_TOKEN_PAIR = errors.New("refresh login error: bad token pair")
)

const (
	SERVICE_NAME = "(RefreshLoginService)"
)

type (
	RotateSignaturesService struct {
		RefreshTokenManipulator    refreshTokenServicePort.IRefreshTokenManipulator
		AccessTokenManipulator     accessTokenServicePort.IAccessTokenManipulator
		AuthSignatureTokenProvider authSignatureTokenPort.IAuthSignatureProvider
		unitOfWork.OperationLogger
	}
)

func (this *RotateSignaturesService) Serve(
	accessToken accessTokenServicePort.IAccessToken, refreshToken refreshTokenServicePort.IRefreshToken, reqCtx context.Context,
) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error) {

	if !accessToken.Expired() {

		err = errors.Join(
			common.ERR_BAD_REQUEST,
			fmt.Errorf("(RefreshLoginService) access token not expired"),
		)
		return
	}

	switch {
	case refreshToken.Expired():
		err = errors.Join(
			common.ERR_UNAUTHORIZED,
			fmt.Errorf("(RefreshLoginService) refresh Token expired"),
		)
		return
	case refreshToken.GetTokenID() != accessToken.GetTokenID(), refreshToken.GetUserUUID() != accessToken.GetUserUUID():
		err = errors.Join(
			common.ERR_UNAUTHORIZED,
			fmt.Errorf("(RefreshLoginService) signatures mismatch"),
		)
		return
	}

	// refreshTokenRevoked, err := this.RefreshTokenBlackList.Has(refreshToken.GetTokenID(), reqCtx)

	// switch {
	// case err != nil:
	// 	return
	// case refreshTokenRevoked:
	// 	err = errors.Join(
	// 		common.ERR_UNAUTHORIZED,
	// 		fmt.Errorf("(RefreshLoginService) refresh token revoked"),
	// 	) //authServiceAdapter.ERR_REFRESH_TOKEN_EXPIRE
	// 	return
	// 	// case refreshToken.GetUserUUID() != accessToken.GetUserUUID():
	// 	// 	err = errors.New() //authServiceAdapter.ERR_REFESH_LOGIN_INVALID_ACCESS_TOKEN
	// 	// 	return
	// }

	return this.AuthSignatureTokenProvider.Rotate(refreshToken, accessToken, reqCtx)
}
