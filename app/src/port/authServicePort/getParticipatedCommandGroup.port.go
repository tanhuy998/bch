package authServicePort

import (
	"app/src/model"
	"app/src/valueObject"
	"context"
)

type (
	IGetParticipatedCommandGroups interface {
		Serve(userUUID string) (*valueObject.ParticipatedCommandGroupReport, error)
		SearchAndRetrieveByModel(model *model.User, ctx context.Context) (*valueObject.ParticipatedCommandGroupReport, error)
	}
)
