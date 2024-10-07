package authServicePort

import (
	"app/port/generalTokenServicePort"
	"context"
)

type (
	ILogIn interface {
		Serve(username string, password string, ctx context.Context) (generalToken generalTokenServicePort.IGeneralToken, err error)
	}
)
