package generalTokenServicePort

import (
	"app/internal/generalToken"
	noExpireTokenServicePort "app/port/noExpireToken"
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	IGeneralToken interface {
		GetUserUUID() uuid.UUID
		GetExpiretime() *time.Time
		GetTokenID() generalToken.GeneralTokenID
	}

	IGeneralTokenProvider interface {
		noExpireTokenServicePort.INoExpireTokenProvider
		Generate(
			userUUID uuid.UUID, ctx context.Context,
		) (IGeneralToken, error)
		GetDefaultExpireDuration() time.Duration
	}

	IGeneralTokenReader interface {
		Read(str string) (IGeneralToken, error)
	}

	IGeneralTokenSigning interface {
		SignString(IGeneralToken) (string, error)
	}

	IGeneralTokenManipulator interface {
		IGeneralTokenProvider
		IGeneralTokenReader
		IGeneralTokenSigning
	}
)
