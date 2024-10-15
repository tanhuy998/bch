package accessTokenServicePort

import (
	"app/valueObject"

	generalTokenServicePort "app/port/generalToken"
	noExpireTokenServicePort "app/port/noExpireToken"
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	AccessTokenAudience struct {
	}

	/*
		IAccessTokenProvider
	*/
	IAccessTokenProvider interface {
		//GenerateByCredentials(model *model.User, tokenID string, ctx context.Context) (IAccessToken, error)
		//GenerateByUserUUID(userUUID uuid.UUID, tokenID string, ctx context.Context) (IAccessToken, error)
		noExpireTokenServicePort.INoExpireTokenProvider
		GenerateBased(IAccessToken IAccessToken, ctx context.Context) (IAccessToken, error)

		GenerateFor(
			tenantUUID uuid.UUID, generalToken generalTokenServicePort.IGeneralToken, tokenID string, ctx context.Context,
		) (IAccessToken, error)
		DefaultExpireDuration() time.Duration

		// CtxNoExpireKey() string
		// IsNoExpire(ctx context.Context) bool
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

	// IParticipatedCommandGroup interface {
	// 	GetCommandGroupUUID() *uuid.UUID
	// 	GetCommandGroupRoleName() string
	// 	HasRoles(name ...string) bool
	// }

	IParticipatedCommandGroup = valueObject.IParticipatedCommandGroup

	// IAccessTokenAuthData interface {
	// 	GetUserUUID() uuid.UUID
	// 	GetTenantUUID() uuid.UUID
	// 	GetTenantAgentData() *model.TenantAgent
	// 	GetParticipatedGroups() []IParticipatedCommandGroup
	// 	IsTenantAgent() bool
	// }

	IAccessTokenAuthData = valueObject.IAuthorityData

	IAccessToken interface {
		// GetExpirationTime() time.Time
		// GetIssuedAt() time.Time
		// GetNotBefore() time.Time
		// GetIssuer() string
		// GetSubject() *uuid.UUID
		// GetAudience() *AccessTokenAudience

		GetTenantUUID() uuid.UUID
		GetUserUUID() uuid.UUID
		GetAudiences() []string
		GetAuthData() IAccessTokenAuthData
		Expired() bool
		GetExpire() *time.Time
		GetTokenID() string
		SetTokenID(string)
	}

	IAccessTokenBuilder interface {
		Build() IAccessToken
	}
)
