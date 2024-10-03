package authServicePort

import (
	"app/model"
	"context"
)

type (
	ICreateCommandGroup interface {
		Serve(groupName string, ctx context.Context) error
		CreateByModel(model *model.CommandGroup, ctx context.Context) (*model.CommandGroup, error)
	}
)
