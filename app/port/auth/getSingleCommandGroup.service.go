package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleCommandGroup interface {
		Serve(uuid uuid.UUID, ctx context.Context) (*model.CommandGroup, error)
		SearchByName(groupName string, ctx context.Context) (*model.CommandGroup, error)
		CheckCommandGroupExistence(groupName string, ctx context.Context) (bool, error)
	}
)
