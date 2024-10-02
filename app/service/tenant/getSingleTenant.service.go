package tenantService

import (
	"app/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleTenant interface {
		Serve(uuid string) (*model.Tenant, error)
	}

	GetSingleTenantService struct {
		TenantRepo repository.ITenant
	}
)

func (this *GetSingleTenantService) Serve(uuid_str string) (*model.Tenant, error) {

	uuid, err := uuid.Parse(uuid_str)

	if err != nil {

		return nil, err
	}

	return this.TenantRepo.FindOneByUUID(uuid, context.TODO())
}
