package generalTokenClientServicePort

import (
	"app/port/generalTokenServicePort"
	"context"
)

type (
	IGeneralTokenClient interface {
		Read(ctx context.Context) (generalTokenServicePort.IGeneralToken, error)
		Write(ctx context.Context, generalToken generalTokenServicePort.IGeneralToken) error
	}
)
