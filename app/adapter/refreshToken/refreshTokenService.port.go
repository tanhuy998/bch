package refreshtokenServicePort

import (
	"context"

	"github.com/google/uuid"
)

type (
	RefreshTokenID = []byte

	IRefreshToken interface {
		GetUserUUID() uuid.UUID
		GetTokenID() string
	}

	IRefreshTokenProvider interface {
		Generate(userUUID uuid.UUID, ctx context.Context) (IRefreshToken, error)
		Revoke(RefreshTokenID) error
	}

	IRefreshTokenSigning interface {
		SignString(IRefreshToken) (string, error)
	}

	IRefreshTokenReader interface {
		Read(string) (IRefreshToken, error)
	}

	IRefreshTokenHadler interface {
		IRefreshTokenProvider
		IRefreshTokenReader
		IRefreshTokenSigning
	}
)
