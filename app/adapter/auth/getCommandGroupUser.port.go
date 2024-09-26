package authServiceAdapter

import (
	"app/domain/model"
	"context"
)

type (
	IGetCommandGroupUsers interface {
		SearchAndRetrieveByModel(dataModel *model.CommandGroup, ctx context.Context) ([]*model.User, error)
		Serve(groupUUID string) ([]*model.User, error)
	}
)
