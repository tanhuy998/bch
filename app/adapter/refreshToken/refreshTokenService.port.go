package refreshTokenServicePort

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ERR_TOKEN_EXPIRE             = errors.New("refresh token expired")
	ERR_REFRESH_TOKEN_BLACK_LIST = errors.New("refresh login error: refresh token in black list")
)

type (
	IRefreshToken interface {
		Expired() bool
		GetUserUUID() uuid.UUID
		GetTokenID() string
		GetExpireTime() (*time.Time, error)
	}

	IRefreshTokenProvider interface {
		Generate(userUUID uuid.UUID, ctx context.Context) (IRefreshToken, error)
		Revoke(refreshToken IRefreshToken, ctx context.Context) error
		DefaultExpireDuration() time.Duration
	}

	IRefreshTokenSigning interface {
		SignString(IRefreshToken) (string, error)
	}

	IRefreshTokenReader interface {
		Read(string) (IRefreshToken, error)
	}

	IRefreshTokenRotation interface {
		Rotate(refreshToken IRefreshToken, ctx context.Context) (IRefreshToken, error)
	}

	IRefreshTokenManipulator interface {
		IRefreshTokenProvider
		IRefreshTokenReader
		IRefreshTokenSigning
		IRefreshTokenRotation
	}
)
