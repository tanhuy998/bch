package assignmentService

import (
	"app/domain/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
)

type (
	GetSingleAssignmentGroupService struct {
		AssignmentGroupRepo repository.IAssignmentGroup
	}
)

func (this *GetSingleAssignmentGroupService) Serve(uuid uuid.UUID, ctx context.Context) (*model.AssignmentGroup, error) {

	return this.AssignmentGroupRepo.FindOneByUUID(uuid, ctx)
}
