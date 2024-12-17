package getAssignmentsDomain

import (
	"app/domain"
	"app/model"
	assignmentServicePort "app/port/assignment"
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
		unitOfWork.PaginateUseCase[model.Assignment, primitive.ObjectID]
	}
)

func (this *GetAssignmentsService) Serve(
	tenantUUID uuid.UUID, filter assignmentServicePort.IGetAssignmentPaginate[primitive.ObjectID], ctx context.Context,
) ([]model.Assignment, error) {

	// var (
	// 	retrieveOnlyExpiredAssignments = filter.GetExpiredFilter()
	// )

	// return this.Paginate(
	// 	tenantUUID,
	// 	ctx,
	// 	paginateUseCaseOption.ByCursor(filter.GetCursor()),
	// 	paginateUseCaseOption.ByOffsetWhenNoCursor(filter.GetPageNumber(), filter.GetPageSize()),
	// 	paginateUseCaseOption.Filter(
	// 		func(filter repositoryAPI.IFilterGenerator) {

	// 			if retrieveOnlyExpiredAssignments {

	// 				filter.Field("deadline").LessThan(time.Now())
	// 			}
	// 		},
	// 	),
	// )

	return this.UseCustomPaginator(
		tenantUUID, filter, ctx,
	)
}
