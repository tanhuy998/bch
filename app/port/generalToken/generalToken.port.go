package generalTokenServicePort

import (
	"app/internal/generalToken"
	noExpireTokenServicePort "app/port/noExpireToken"
	jwtClaim "app/valueObject/jwt"
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	IGeneralToken interface {
		GetPolicies() []jwtClaim.GenTokenPolicyEnum
		GetUserUUID() uuid.UUID
		GetExpiretime() *time.Time
		GetTokenID() generalToken.GeneralTokenID
		Expire() bool
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
