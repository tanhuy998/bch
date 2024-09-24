package tenantServicePort

import (
	"app/domain/model"
	"context"
)

type (
	ICreateTenant interface {
		Serve(tenantInput *model.Tenant, userInput *model.User, ctx context.Context) (*model.Tenant, *model.User, error)
	}
)
