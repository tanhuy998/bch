package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetCommandGroupUsers interface {
		Serve(groupUUID uuid.UUID, ctx context.Context) ([]*model.User, error)
	}
)
