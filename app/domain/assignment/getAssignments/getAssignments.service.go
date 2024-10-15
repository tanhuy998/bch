package getAssignmentsDomain

import (
	"app/domain"
	libCommon "app/internal/lib/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	"app/repository"
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	DEFAULT_PAGE_SIZE = 10
)

type (
	GetAssignmentsService struct {
		domain.ContextualDomainService[domain_context]
		AssignmentRepo repository.IAssignment
	}
)

func (this *GetAssignmentsService) Serve(
	TenantUUID uuid.UUID, filter assignmentServicePort.IGetAssignmentPaginate, ctx context.Context,
) ([]model.Assignment, error) {

	var (
		size       = filter.GetPageSize()
		pageNumber = filter.GetPageNumber()
		expired    = filter.GetExpiredFilter()
	)

	if size <= 0 {

		size = DEFAULT_PAGE_SIZE
	}

	var query bson.D

	switch expired {
	case true:
		query = bson.D{
			{"tenantUUID", TenantUUID},
			{
				"deadline", bson.D{
					{"$lt", time.Now()},
				},
			},
		}
	case false:
		query = bson.D{
			{"tenantUUID", TenantUUID},
		}
	}

	if !filter.HasCursor() {
		return this.AssignmentRepo.FindOffset(
			query, pageNumber, size, nil, ctx,
		)
	}

	model := model.Assignment{}
	model.ObjectID = libCommon.PointerPrimitive(filter.GetCursor())

	switch filter.IsPrevious() {
	case true:
		return repository.FindPrevious(this.AssignmentRepo.GetCollection(), model, size, ctx, query...)
	default:
		return repository.FindNext(this.AssignmentRepo.GetCollection(), model, size, ctx, query...)
	}
}
