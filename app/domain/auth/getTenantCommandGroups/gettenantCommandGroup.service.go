package getTenantCommandGroupDomain

import (
	"app/domain"
	"app/model"
	"app/repository"
	repositoryAPI "app/repository/api"
	"app/unitOfWork"
	"context"

	"github.com/google/uuid"
)

type (
	GetTenantCommandGroupService struct {
		unitOfWork.PaginateUseCase[
			repository.ICommandGroup,
			model.CommandGroup,
			interface{},
		]
	}
)

func (this *GetTenantCommandGroupService) Serve(
	tenantUUID uuid.UUID, paginator domain.IPaginator, ctx context.Context,
) ([]model.CommandGroup, error) {

	switch {
	case paginator == nil:
		return this.ServeWithoutPaginator(
			tenantUUID, ctx,
		)
	default:
		return this.PaginateUseCase.UseCustomPaginator(
			tenantUUID, paginator, ctx,
		)
	}
}

func (this *GetTenantCommandGroupService) ServeWithoutPaginator(
	tenantUUID uuid.UUID, ctx context.Context,
) ([]model.CommandGroup, error) {

	return this.Repository.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	).Find(ctx)
}
