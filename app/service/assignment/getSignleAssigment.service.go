package assignmentService

import (
	"app/domain/model"
	"app/repository"
	"context"

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

	return this.AssignmentRepo.Find(
		bson.D{
			{"title", model.Title},
			{"tenantUUID", model.TenantUUID},
		},
		ctx,
	)
}
