package accessTokenServicePort

import (
	"app/domain/model"
	"context"

	"github.com/google/uuid"
)

type (
	AccessTokenAudience struct {
	}

	IAccessTokenProvider interface {
		Generate(userUUID uuid.UUID, ctx context.Context) (IAccessToken, error)
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

		GetUserUUID() *uuid.UUID
		GetAudiences() []string
		GetAuthData() IAccessTokenAuthData
		Expired() bool
	}

	IAccessTokenBuilder interface {
		Build() IAccessToken
	}
)
