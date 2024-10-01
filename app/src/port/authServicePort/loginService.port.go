package authServicePort

import (
	"context"
)

type (
	ILogIn interface {
		Serve(username string, password string, ctx context.Context) (accessToken string, refreshToken string, err error)
	}
)
