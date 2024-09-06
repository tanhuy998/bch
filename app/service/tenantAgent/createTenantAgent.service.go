package tenantAgentService

import (
	authServiceAdapter "app/adapter/auth"
	passwordServiceAdapter "app/adapter/passwordService"
	"app/domain/model"
	libCommon "app/lib/common"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_TENANT_AGENT_EXISTS = errors.New("tenant agent exists")
)

type (
	ICreaateTenantAgent interface {
		//Serve(dataModel *model.User) (*model.TenantAgent, error)
		Serve(username string, password string, name string) (*model.TenantAgent, error)
	}

	CreateTenantAgentService struct {
		GetSingleTenantService IGetSingleTenantAgent
		GetSingleUserService   authServiceAdapter.IGetSingleUserService
		CreateUserService      authServiceAdapter.ICreateUserService
		TenantAgentRepo        repository.ITenantAgent
		PasswordService        passwordServiceAdapter.IPassword
	}
)

func (this CreateTenantAgentService) Serve(username string, password string, name string) (*model.TenantAgent, error) {

	newUser, err := this.CreateUserService.Serve(username, password, name)

	if err != nil {

		return nil, err
	}

	newAgentModel := &model.TenantAgent{
		UUID:        libCommon.PointerPrimitive(uuid.New()),
		UserUUID:    &newUser.UUID,
		Deactivated: true,
	}

	err = this.TenantAgentRepo.Create(newAgentModel, context.TODO())

	if err != nil {

		return nil, err
	}

	return this.GetSingleTenantService.Serve(newAgentModel.UUID.String())
}
