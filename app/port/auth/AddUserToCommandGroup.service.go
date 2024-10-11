package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IAddUserToCommandGroup interface {
		Serve(tenantUUID uuid.UUID, dataModel *model.CommandGroupUser, ctx context.Context) error
		Get() IGetSingleCommandGroup
	}
)
