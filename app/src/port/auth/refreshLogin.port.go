package authServiceAdapter

import (
	"context"
	"errors"
)

var (
	ERR_ACCESS_TOKEN_NOT_EXPIRE           = errors.New("refresh login error: access token not expire")
	ERR_REFRESH_TOKEN_EXPIRE              = errors.New("refresh login error: refresh token expire, need reautherization.")
	ERR_REFESH_LOGIN_INVALID_ACCESS_TOKEN = errors.New("refresh login error: invalic access token")
)

type (
	IRefreshLogin interface {
		Serve(inputAT string, inputRT string, reqCtx context.Context) (at string, rt string, err error)
	}
)
