package authGenServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	NavigateTenantInput interface {
		GetRequestedUserUUID() uuid.UUID
		GetContext() context.Context
	}

	INavigateTenant interface {
		//Serve(userUUID uuid.UUID, ctx context.Context) ([]model.Tenant, error)
		Serve(input NavigateTenantInput) ([]model.Tenant, error)
	}
)
