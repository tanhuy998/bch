package getSingleAssignmentDomain

import (
	"app/internal/common"
	"app/model"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	GetSingleAssignmentService struct {
		AssignmentRepo repository.IAssignment
	}
)

func (this *GetSingleAssignmentService) Serve(
	uuid uuid.UUID, ctx context.Context,
) (*model.Assignment, error) {

	return this.AssignmentRepo.FindOneByUUID(uuid, ctx)
}

func (this *GetSingleAssignmentService) Search(
	model *model.Assignment, ctx context.Context,
) (*model.Assignment, error) {

	ret, err := this.AssignmentRepo.Find(
		bson.D{
			{"title", model.Title},
			{"tenantUUID", model.TenantUUID},
		},
		ctx,
	)

	if err != nil {

		return nil, errors.Join(common.ERR_INTERNAL, err)
	}

	return ret, nil
}
