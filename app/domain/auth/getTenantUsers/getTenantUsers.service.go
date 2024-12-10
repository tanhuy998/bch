package getTenantUsersDomain

import (
	"app/model"
	"app/unitOfWork"
	paginateUseCaseOption "app/unitOfWork/genericUsecase/paginate/option"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetTenantUsersService struct {
		unitOfWork.PaginateUseCase[model.User]
	}
)

func (this *GetTenantUsersService) Serve(
	tenantUUID uuid.UUID, page uint64, size uint64, cursor *primitive.ObjectID, isPrev bool, ctx context.Context,
) ([]model.User, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	return this.Paginate(
		tenantUUID, ctx,
		paginateUseCaseOption.ByCursor(cursor),
		paginateUseCaseOption.ByOffsetWhenNoCursor(page, size),
	)
}
