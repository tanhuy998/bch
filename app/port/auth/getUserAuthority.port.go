package authServicePort

import (
	"app/valueObject"
	"context"

	"github.com/google/uuid"
)

type (
	IGetUserAuthorityServicePort interface {
		Serve(
			tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
		) (*valueObject.AuthData, error)
	}
)
