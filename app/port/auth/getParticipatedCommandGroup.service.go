package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetUserParticipatedCommandGroups interface {
		Serve(
			tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
		) ([]*model.CommandGroup, error)
		SearchAndRetrieveByModel(
			searchModel *model.User, ctx context.Context,
		) ([]*model.CommandGroup, error)
	}
)
