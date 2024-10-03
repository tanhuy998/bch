package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IModifyUser interface {
		Serve(userUUID uuid.UUID, data *model.User, ctx context.Context) error
	}
)
