package getParticipatedCommandGroup

import (
	"app/model"
	"app/valueObject"
	"context"

	"github.com/google/uuid"
)

type (
	GetParticipatedCommandGroupService struct {
	}
)

func (this *GetParticipatedCommandGroupService) Serve(
	userUUID uuid.UUID, ctx context.Context,
) (*valueObject.ParticipatedCommandGroupReport, error) {
	return nil, nil
}

func (this *GetParticipatedCommandGroupService) SearchAndRetrieveByModel(
	searchModel *model.User, ctx context.Context,
) (*valueObject.ParticipatedCommandGroupReport, error) {

	return nil, nil
}
