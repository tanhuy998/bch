package accessTokenServicePort

import (
	"app/domain/model"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	AccessTokenAudience struct {
	}

	IAccessTokenProvider interface {
		GenerateByCredentials(model *model.User, tokenID string, ctx context.Context) (IAccessToken, error)
		GenerateByUserUUID(userUUID uuid.UUID, tokenID string, ctx context.Context) (IAccessToken, error)
		GenerateBased(IAccessToken IAccessToken, ctx context.Context) (IAccessToken, error)
		DefaultExpireDuration() time.Duration
	}

	IAccessTokenSigning interface {
		SignString(IAccessToken) (string, error)
	}

	IAccessTokenReader interface {
		Read(string) (IAccessToken, error)
	}

	AccessTokenAudienceList []string

	// IAccessTokenRefresh interface {
	// 	Refresh(string) (IAccessToken, error)
	// }

	IAccessTokenManipulator interface {
		IAccessTokenProvider
		IAccessTokenReader
		IAccessTokenSigning
	}

	IParticipatedCommandGroup interface {
		GetCommandGroupUUID() *uuid.UUID
		GetCommandGroupRoleName() string
	}

	IAccessTokenAuthData interface {
		GetTenantUUID() *uuid.UUID
		GetTenantAgentData() *model.TenantAgent
		GetParticipatedGroups() []IParticipatedCommandGroup
		IsTenantAgent() bool
	}

	IAccessToken interface {
		// GetExpirationTime() time.Time
		// GetIssuedAt() time.Time
		// GetNotBefore() time.Time
		// GetIssuer() string
		// GetSubject() *uuid.UUID
		// GetAudience() *AccessTokenAudience

		GetUserUUID() uuid.UUID
		GetAudiences() []string
		GetAuthData() IAccessTokenAuthData
		Expired() bool
		GetExpire() (*jwt.NumericDate, error)
		GetTokenID() string
		SetTokenID(string)
	}

	IAccessTokenBuilder interface {
		Build() IAccessToken
	}
)
