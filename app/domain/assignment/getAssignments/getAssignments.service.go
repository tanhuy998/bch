package getAssignmentsDomain

import (
	"app/domain"
	"app/model"
	assignmentServicePort "app/port/assignment"
	repositoryAPI "app/repository/api"
	"app/unitOfWork"
	paginateUseCaseOption "app/unitOfWork/genericUsecase/paginate/option"
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	DEFAULT_PAGE_SIZE = 10
)

type (
	GetAssignmentsService struct {
		domain.ContextualDomainService[domain_context]
		unitOfWork.PaginateUseCase[model.Assignment]
	}
)

func (this *GetAssignmentsService) Serve(
	tenantUUID uuid.UUID, filter assignmentServicePort.IGetAssignmentPaginate, ctx context.Context,
) ([]model.Assignment, error) {

	var (
		expired = filter.GetExpiredFilter()
	)

	return this.Paginate(
		tenantUUID,
		ctx,
		paginateUseCaseOption.ByCursor(filter.GetCursor()),
		paginateUseCaseOption.ByOffsetWhenNoCursor(filter.GetPageNumber(), filter.GetPageSize()),
		paginateUseCaseOption.Filter(
			func(filter repositoryAPI.IFilterGenerator) {

				if expired {

					filter.Field("deadline").LessThan(time.Now())
				}
			},
		),
	)
}
