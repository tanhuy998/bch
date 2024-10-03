package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	ICheckUserInCommandGroup interface {
		Serve(groupUUID, userUUID uuid.UUID, ctx context.Context) (bool, error)
		Detail(groupUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context) (*model.CommandGroupUser, error)
	}
)
