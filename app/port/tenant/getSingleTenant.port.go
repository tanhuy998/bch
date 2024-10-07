package tenantServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleTenant interface {
		Serve(uuid uuid.UUID, ctx context.Context) (*model.Tenant, error)
		CheckExist(uuit uuid.UUID, ctx context.Context) (bool, error)
	}
)
