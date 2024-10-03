package authServicePort

import (
	"context"

	"github.com/google/uuid"
)

type (
	IAddUserToCommandGroup interface {
		Serve(groupUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context) error
		Get() IGetSingleCommandGroup
	}
)
