package tenantAgentService

import (
	passwordServiceAdapter "app/adapter/passwordService"
	"app/domain/model"
	"app/repository"
	"context"
	"errors"
)

var (
	ERR_TENANT_AGENT_EXISTS = errors.New("tenant agent exists")
)

type (
	ICreaateTenantAgent interface {
		Serve(dataModel *model.TenantAgent) error
	}

	CreateTenantAgentService struct {
		GetSingleTenantService IGetSingleTenantAgent
		TenantAgentRepo        repository.ITenantAgent
		PasswordService        passwordServiceAdapter.IPassword
	}
)

func (this CreateTenantAgentService) Serve(dataModel *model.TenantAgent) error {

	err := this.PasswordService.Resolve(dataModel)

	if err != nil {

		return err
	}

	existsTenant, err := this.GetSingleTenantService.SearchByUsername(dataModel.Username)

	if err != nil {

		return err
	}

	if existsTenant != nil {

		return ERR_TENANT_AGENT_EXISTS
	}

	err = this.TenantAgentRepo.Create(dataModel, context.TODO())

	if err != nil {

		return err
	}

	return nil
}
