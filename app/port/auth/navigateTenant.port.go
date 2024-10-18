package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	INavigateTenant interface {
		Serve(userUUID uuid.UUID, ctx context.Context) ([]*model.Tenant, error)
	}
)
