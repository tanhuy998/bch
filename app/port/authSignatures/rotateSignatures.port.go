package authSignaturesServicePort

import (
	accessTokenServicePort "app/port/accessToken"
	refreshTokenServicePort "app/port/refreshToken"
	"context"
	"errors"
)

var (
	ERR_ACCESS_TOKEN_NOT_EXPIRE           = errors.New("refresh login error: access token not expire")
	ERR_REFRESH_TOKEN_EXPIRE              = errors.New("refresh login error: refresh token expire, need reautherization.")
	ERR_REFESH_LOGIN_INVALID_ACCESS_TOKEN = errors.New("refresh login error: invalic access token")
)

type (
	IRotateSignatures interface {
		Serve(
			inputAT accessTokenServicePort.IAccessToken, inputRT refreshTokenServicePort.IRefreshToken, reqCtx context.Context,
		) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error)
	}
)
