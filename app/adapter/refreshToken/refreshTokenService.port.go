package refreshTokenServicePort

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	IRefreshToken interface {
		GetUserUUID() uuid.UUID
		GetTokenID() string
	}

	IRefreshTokenProvider interface {
		Generate(userUUID uuid.UUID, ctx context.Context) (IRefreshToken, error)
		Revoke(tokenID string, ctx context.Context) error
		DefaultExpireDuration() time.Duration
	}

	IRefreshTokenSigning interface {
		SignString(IRefreshToken) (string, error)
	}

	IRefreshTokenReader interface {
		Read(string) (IRefreshToken, error)
	}

	IRefreshTokenManipulator interface {
		IRefreshTokenProvider
		IRefreshTokenReader
		IRefreshTokenSigning
	}
)
