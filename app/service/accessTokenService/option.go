package accessTokenService

import "time"

type (
	AccessTokenManipulatorOption func(*JWTAccessTokenManipulatorService)
)

func WithoutExpire(m *JWTAccessTokenManipulatorService) {

	m.WithoutExpire = true
}

func ExpireDuration(d time.Duration) AccessTokenManipulatorOption {

	return func(m *JWTAccessTokenManipulatorService) {

		m.ExpDuration = d
	}
}
