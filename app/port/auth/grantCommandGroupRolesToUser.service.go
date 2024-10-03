package authServicePort

import (
	"context"

	"github.com/google/uuid"
)

type (
	IGrantCommandGroupRolesToUser interface {
		Serve(groupUUID uuid.UUID, userUUID uuid.UUID, roles []uuid.UUID, ctx context.Context) error
	}
)
