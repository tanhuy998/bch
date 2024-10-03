package authServicePort

import (
	"app/model"
	"context"
)

type (
	IGetAllRoles interface {
		Serve(ctx context.Context) ([]*model.Role, error)
	}
)
