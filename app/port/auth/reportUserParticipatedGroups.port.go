package authServicePort

import (
	"app/model"
	"app/valueObject"
	"context"

	"github.com/google/uuid"
)

type (
	IReportParticipatedCommandGroups interface {
		Serve(
			tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
		) (*valueObject.ParticipatedCommandGroupReport, error)
		SearchAndRetrieveByModel(
			searchModel *model.User, ctx context.Context,
		) (*valueObject.ParticipatedCommandGroupReport, error)
	}
)
