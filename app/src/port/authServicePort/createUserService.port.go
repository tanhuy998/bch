package authServicePort

import (
	"app/src/model"
	"context"
)

type (
	ICreateUserService interface {
		Serve(username string, password string, name string, ctx context.Context) (*model.User, error)
		CreateByModel(dataModel *model.User, ctx context.Context) (*model.User, error)
	}
)
