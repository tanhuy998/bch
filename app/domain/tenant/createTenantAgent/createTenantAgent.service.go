package createTenantAgentDomain

import (
	libError "app/internal/lib/error"
	"app/model"
	authServicePort "app/port/auth"
	passwordServicePort "app/port/passwordService"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
		// GetSingleUserService authServicePort.IGetSingleUserService
		CreateUserService authServicePort.ICreateUser
		TenantAgentRepo   repository.ITenantAgent
		PasswordService   passwordServicePort.IPassword
		MongoClient       *mongo.Client
	}
)

func (this CreateTenantAgentService) Serve(inputUser *model.User, newTenantAgent *model.TenantAgent, tenantUUID uuid.UUID, ctx context.Context) (*model.User, *model.TenantAgent, error) {

	if ctx != nil {

		ctx = context.Background()
	}

	inputUser.TenantUUID = &tenantUUID

	session, err := this.MongoClient.StartSession()

	if err != nil {

		return nil, nil, libError.NewInternal(err)
	}

	defer session.EndSession(ctx)

	_, err = session.WithTransaction(
		ctx,
		func(ctx mongo.SessionContext) (interface{}, error) {

			_, err := this.CreateUserService.CreateByModel(inputUser, ctx)

			if err != nil {

				return nil, err
			}

			err = this.TenantAgentRepo.Create(newTenantAgent, ctx)

			if err != nil {

				return nil, err
			}

			return nil, nil
		},
	)

	if err != nil {

		return nil, nil, err
	}
	//return this.GetSingleTenantService.Serve(newAgentModel.UUID.String())
	return inputUser, newTenantAgent, nil
}
