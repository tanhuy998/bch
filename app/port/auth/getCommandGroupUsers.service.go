package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetCommandGroupUsers interface {
		Serve(tenantUUID uuid.UUID, groupUUID uuid.UUID, ctx context.Context) ([]*model.User, error)
	}
)
