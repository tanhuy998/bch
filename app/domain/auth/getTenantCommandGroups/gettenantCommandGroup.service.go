package getTenantCommandGroupDomain

import (
	"app/domain"
	"app/model"
	"app/repository"
	"app/unitOfWork"
	"context"

	"github.com/google/uuid"
)

type (
	GetTenantCommandGroupService struct {
		unitOfWork.PaginateUseCase[repository.ICommandGroup, model.CommandGroup, domain.PaginateCursorType]
	}
)

func (this *GetTenantCommandGroupService) Serve(
	tenantUUID uuid.UUID, pagiantor domain.IPaginator, ctx context.Context,
) ([]model.CommandGroup, error) {

	return this.PaginateUseCase.UseCustomPaginator(
		tenantUUID, pagiantor, ctx,
	)
}
