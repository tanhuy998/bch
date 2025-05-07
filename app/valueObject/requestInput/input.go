package requestInput

import (
	"app/valueObject"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IContextBringAlong interface {
		ReceiveContext(ctx context.Context)
		GetContext() context.Context
	}

	IAuthorityGetter interface {
		GetAuthority() valueObject.IAuthorityData
	}

	IAuthoritySetter interface {
		SetAuthority(auth valueObject.IAuthorityData)
	}

	IAuthorityBringAlong interface {
		IAuthorityGetter
		IAuthoritySetter
	}

	ITenantMappingInput interface {
		SetTenantUUID(tenantUUID uuid.UUID)
		IsValidTenantUUID() bool
		GetTenantUUID() uuid.UUID
	}

	IPaginationInput interface {
		GetPageNumber() uint64
		GetPageSize() uint64
	}

	ICursorPaginationInput[Cursor_Type comparable] interface {
		GetCursor() Cursor_Type
		HasCursor() bool
		IsPrevious() bool
	}

	IMongoCursorPaginationInput interface {
		ICursorPaginationInput[primitive.ObjectID]
	}

	ITenantDomainInput interface {
		IContextBringAlong
		IAuthorityBringAlong
		ITenantMappingInput
	}
)
