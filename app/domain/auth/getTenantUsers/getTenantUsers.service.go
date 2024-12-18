package getTenantUsersDomain

import (
	"app/model"
	paginateServicePort "app/port/paginate"
	"app/repository"
	"app/unitOfWork"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetTenantUsersService struct {
		unitOfWork.PaginateUseCase[repository.IUser, model.User, primitive.ObjectID]
	}
)

func (this *GetTenantUsersService) Serve(
	tenantUUID uuid.UUID, paginator paginateServicePort.IPaginator[primitive.ObjectID], ctx context.Context,
) ([]model.User, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	return this.PaginateUseCase.UseCustomPaginator(
		tenantUUID, paginator, ctx,
	)
}
