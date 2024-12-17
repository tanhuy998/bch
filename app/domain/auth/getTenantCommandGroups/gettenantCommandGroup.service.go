package getTenantCommandGroupDomain

import (
	"app/model"
	paginateServicePort "app/port/paginate"
	"app/unitOfWork"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetTenantCommandGroupService struct {
		unitOfWork.PaginateUseCase[model.CommandGroup, primitive.ObjectID]
	}
)

func (this *GetTenantCommandGroupService) Serve(
	//tenantUUID uuid.UUID, page uint64, size uint64, cursor *primitive.ObjectID, isPrev bool, ctx context.Context,
	tenantUUID uuid.UUID, pagiantor paginateServicePort.IPaginator[primitive.ObjectID], ctx context.Context,
) ([]model.CommandGroup, error) {

	// return this.PaginateUseCase.Paginate(
	// 	tenantUUID, ctx,
	// 	paginateUseCaseOption.ByCursor(cursor),
	// 	paginateUseCaseOption.ByOffsetWhenNoCursor(page, size),
	// )

	return this.PaginateUseCase.UseCustomPaginator(
		tenantUUID, pagiantor, ctx,
	)
}
