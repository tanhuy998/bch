package authServicePort

import (
	"app/model"
	paginateServicePort "app/port/paginate"
	"context"

	"github.com/google/uuid"
)

type (
	GetTenantAllGroupsInput interface {
		GetTenantUUID() uuid.UUID
		//paginateServicePort.IGeneralPaginator
		GetPaginator() paginateServicePort.IPaginator[interface{}]
		GetContext() context.Context
	}

	IGetTenantAllGroups interface {
		//Serve(tenantUUID uuid.UUID, ctx context.Context) ([]*model.CommandGroup, error)
		Serve(input GetTenantAllGroupsInput) ([]model.CommandGroup, error)
	}
)
