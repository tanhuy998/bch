package tenantService

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
		Serve(inputUser *model.User, tenantUUID uuid.UUID, ctx context.Context) (*model.User, *model.TenantAgent, error)
	}

	CreateTenantAgentService struct {
		//GetSingleTenantService IGetSingleTenantAgent
		GetSingleUserService authServiceAdapter.IGetSingleUserService
		CreateUserService    authServiceAdapter.ICreateUserService
		TenantAgentRepo      repository.ITenantAgent
		PasswordService      passwordServiceAdapter.IPassword
	}
)

func (this CreateTenantAgentService) Serve(inputUser *model.User, tenantUUID uuid.UUID, ctx context.Context) (*model.User, *model.TenantAgent, error) {

	inputUser.TenantUUID = &tenantUUID

	newUser, err := this.CreateUserService.CreateByModel(inputUser, ctx)

	if err != nil {

		return nil, nil, err
	}

	newAgentModel := &model.TenantAgent{
		UUID:       libCommon.PointerPrimitive(uuid.New()),
		UserUUID:   libCommon.PointerPrimitive(uuid.UUID(*newUser.UUID)),
		TenantUUID: libCommon.PointerPrimitive(tenantUUID),
		//Deactivated: true,
	}

	err = this.TenantAgentRepo.Create(newAgentModel, ctx)

	if err != nil {

		return nil, nil, err
	}

	//return this.GetSingleTenantService.Serve(newAgentModel.UUID.String())
	return newUser, newAgentModel, nil
}
