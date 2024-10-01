package tenantServicePort

import (
	"app/src/model"
	"context"

	"github.com/google/uuid"
)

type (
	ICreateTenantAgent interface {
		Serve(inputUser *model.User, tenantUUID uuid.UUID, ctx context.Context) (*model.User, *model.TenantAgent, error)
	}
)
