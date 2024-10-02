package authServicePort

import (
	"app/model"
	"context"
)

type (
	IGetSingleUser interface {
		Serve(uuid string, ctx context.Context) (*model.User, error)
		SearchByUsername(username string, ctx context.Context) (*model.User, error)
		CheckUsernameExistence(username string, ctx context.Context) (bool, error)
	}
)
