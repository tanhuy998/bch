package tenantServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGrantUserAsTenantAgent interface {
		Serve(
			userUUID uuid.UUID, tenantUUID uuid.UUID, newTenantAgent *model.TenantAgent, ctx context.Context,
		) (*model.TenantAgent, error)
	}
)
