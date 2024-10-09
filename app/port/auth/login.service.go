package authServicePort

import (
	generalTokenServicePort "app/port/generalToken"
	"context"
)

type (
	ILogIn interface {
		Serve(username string, password string, ctx context.Context) (generalToken generalTokenServicePort.IGeneralToken, err error)
	}
)
