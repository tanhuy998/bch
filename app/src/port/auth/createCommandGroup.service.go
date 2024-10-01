package authServicePort

import (
	"app/src/model"
	"context"
)

type (
	ICreateCommandGroup interface {
		Serve(groupName string) error
		CreateByModel(model *model.CommandGroup, ctx context.Context) (*model.CommandGroup, error)
	}
)
