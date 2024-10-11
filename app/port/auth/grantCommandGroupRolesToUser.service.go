package authServicePort

import (
	"context"

	"github.com/google/uuid"
)

type (
	IGrantCommandGroupRolesToUser interface {
		Serve(tenantUUID uuid.UUID, groupUUID uuid.UUID, userUUID uuid.UUID, roles []uuid.UUID, createdBy uuid.UUID, ctx context.Context) error
	}
)
