package accessTokenService

import "time"

type (
	AccessTokenManipulatorOption func(*JWTAccessTokenProviderService)
)

func WithoutExpire(m *JWTAccessTokenProviderService) {

	m.WithoutExpire = true
}

func ExpireDuration(d time.Duration) AccessTokenManipulatorOption {

	return func(m *JWTAccessTokenProviderService) {

		m.ExpDuration = d
	}
}
