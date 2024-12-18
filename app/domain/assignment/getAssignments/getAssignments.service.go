package getAssignmentsDomain

import (
	"app/domain"
	"app/model"
	assignmentServicePort "app/port/assignment"
	"app/repository"
	"app/unitOfWork"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DEFAULT_PAGE_SIZE = 10
)

type (
	GetAssignmentsService struct {
		domain.ContextualDomainService[domain_context]
		unitOfWork.PaginateUseCase[repository.IAssignment, model.Assignment, primitive.ObjectID]
	}
)

func (this *GetAssignmentsService) Serve(
	tenantUUID uuid.UUID, filter assignmentServicePort.IGetAssignmentPaginate[primitive.ObjectID], ctx context.Context,
) ([]model.Assignment, error) {

	return this.UseCustomPaginator(
		tenantUUID, filter, ctx,
	)
}
