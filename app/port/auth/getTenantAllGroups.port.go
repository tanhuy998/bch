package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetTenantAllGroups interface {
		Serve(tenantUUID uuid.UUID, ctx context.Context) ([]*model.CommandGroup, error)
	}
)
