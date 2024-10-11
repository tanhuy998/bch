package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	ICreateCommandGroup interface {
		Serve(tenantUUID uuid.UUID, groupName string, ctx context.Context) (*model.CommandGroup, error)
		CreateByModel(tenantUUID uuid.UUID, model *model.CommandGroup, ctx context.Context) (*model.CommandGroup, error)
	}
)
