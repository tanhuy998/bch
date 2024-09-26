package authServiceAdapter

import (
	"app/domain/model"
	"app/domain/valueObject"
	"context"
)

type (
	IGetParticipatedCommandGroups interface {
		Serve(userUUID string) (*valueObject.ParticipatedCommandGroupReport, error)
		SearchAndRetrieveByModel(model *model.User, ctx context.Context) (*valueObject.ParticipatedCommandGroupReport, error)
	}
)
