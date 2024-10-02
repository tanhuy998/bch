package getSingleAssignmentGroupDomain

import (
	"app/internal/common"
	"app/model"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

type (
	GetSingleAssignmentGroupService struct {
		AssignmentGroupRepo repository.IAssignmentGroup
	}
)

func (this *GetSingleAssignmentGroupService) Serve(uuid uuid.UUID, ctx context.Context) (*model.AssignmentGroup, error) {

	ret, err := this.AssignmentGroupRepo.FindOneByUUID(uuid, ctx)

	if err != nil {

		return nil, errors.Join(
			common.ERR_INTERNAL,
			err,
		)
	}

	return ret, nil
}
