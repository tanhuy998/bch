package accessTokenServicePort

import (
	"github.com/google/uuid"
)

type (
	AccessTokenAudience struct {
	}

	IAccessTokenProvider interface {
		Generate(audience *AccessTokenAudience) (IAccessToken, error)
	}

	IAccessTokenSigning interface {
		SignedString(IAccessToken) (string, error)
	}

	IAccessTokenReader interface {
		Read(string) (IAccessToken, error)
	}

	AccessTokenAudienceList []string

	// IAccessTokenRefresh interface {
	// 	Refresh(string) (IAccessToken, error)
	// }

	IAccessTokenHandler interface {
		IAccessTokenProvider
		IAccessTokenReader
		IAccessTokenSigning
	}

	IAccessToken interface {
		// GetExpirationTime() time.Time
		// GetIssuedAt() time.Time
		// GetNotBefore() time.Time
		// GetIssuer() string
		// GetSubject() *uuid.UUID
		// GetAudience() *AccessTokenAudience

		GetUserUUID() *uuid.UUID
		GetAudiences() []string
		Expired() bool
	}

	IAccessTokenBuilder interface {
		Build() IAccessToken
	}
)
