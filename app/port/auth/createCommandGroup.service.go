package authServicePort

import (
	"app/model"
	"context"
)

type (
	ICreateCommandGroup interface {
		Serve(groupName string) error
		CreateByModel(model *model.CommandGroup, ctx context.Context) (*model.CommandGroup, error)
	}
)
