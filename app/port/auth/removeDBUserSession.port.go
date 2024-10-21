package authServicePort

import (
	"context"

	"github.com/google/uuid"
)

type (
	IRemoveDBUserSession interface {
		Serve(userUUID uuid.UUID, ctx context.Context) error
	}
)
