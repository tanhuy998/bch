package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleUser interface {
		Serve(uuid uuid.UUID, ctx context.Context) (*model.User, error)
		SearchByUsername(username string, ctx context.Context) (*model.User, error)
		CheckUsernameExistence(username string, ctx context.Context) (bool, error)
	}
)
