package genericUseCase

import (
	accessTokenServicePort "app/port/accessToken"
	refreshTokenServicePort "app/port/refreshToken"
)

type (
	SignatureManagementUseCase struct {
		RefreshTokenBlackList RefreshTokenBlackList
	}
)

func (this *SignatureManagementUseCase) RevokeSignatures(
	accessToken accessTokenServicePort.IAccessToken, refreshToken refreshTokenServicePort.IRefreshToken,
) error {

	return nil
}
