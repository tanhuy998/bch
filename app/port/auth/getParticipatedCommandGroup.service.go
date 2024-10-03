package authServicePort

import (
	"app/model"
	"app/valueObject"
	"context"

	"github.com/google/uuid"
)

type (
	IGetParticipatedCommandGroups interface {
		Serve(userUUID uuid.UUID, ctx context.Context) (*valueObject.ParticipatedCommandGroupReport, error)
		SearchAndRetrieveByModel(
			searchModel *model.User, ctx context.Context,
		) (*valueObject.ParticipatedCommandGroupReport, error)
	}
)
