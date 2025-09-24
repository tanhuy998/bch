package authenticatorServicePort

import (
	accessTokenServicePort "app/port/accessToken"
	"context"
)

type (
	IAuthenticator interface {
		// Validate(ctx context.Context)
		// Verify(ctx context.Context)
		VerifyAccessToken(ctx context.Context) (accessToken accessTokenServicePort.IAccessToken, err error)
		Authorize(ctx context.Context, policies ...IAuthorizePolicy) error
	}
)

type (
	IAuthorizePolicy interface {
		Apply(accessTokenServicePort.IAccessToken) error
	}
)
