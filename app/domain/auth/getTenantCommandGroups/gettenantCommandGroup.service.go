package getTenantCommandGroupDomain

import (
	"app/model"
	"app/unitOfWork"
	paginateUseCaseOption "app/unitOfWork/genericUsecase/paginate/option"

	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetTenantCommandGroupService struct {
		unitOfWork.PaginateUseCase[model.CommandGroup]
	}
)

func (this *GetTenantCommandGroupService) Paginate(
	tenantUUID uuid.UUID, page uint64, size uint64, cursor *primitive.ObjectID, isPrev bool, ctx context.Context,
) ([]model.CommandGroup, error) {

	return this.PaginateUseCase.Paginate(
		tenantUUID, ctx,
		paginateUseCaseOption.ByCursor(cursor),
		paginateUseCaseOption.ByOffsetWhenNoCursor(page, size),
	)
}
