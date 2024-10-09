package getSingleTenantDomain

import (
	"app/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
)

type (
	GetSingleTenantService struct {
		TenantRepo repository.ITenant
	}
)

func (this *GetSingleTenantService) Serve(uuid uuid.UUID, ctx context.Context) (*model.Tenant, error) {

	return this.TenantRepo.FindOneByUUID(uuid, ctx)
}

func (this *GetSingleTenantService) CheckExist(uuid uuid.UUID, ctx context.Context) (bool, error) {

	m, err := this.Serve(uuid, ctx)

	if err != nil {

		return false, err
	}

	return m != nil, nil
}
