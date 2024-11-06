package authGenServicePort

import (
	generalTokenServicePort "app/port/generalToken"
	"context"
)

type (
	IAuthenticateCrdentials interface {
		Serve(username string, password string, ctx context.Context) (generalToken generalTokenServicePort.IGeneralToken, err error)
	}
)
