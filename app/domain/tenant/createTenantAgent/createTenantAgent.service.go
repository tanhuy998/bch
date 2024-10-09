package createTenantAgentDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"app/internal/lib/libContext"
	"app/model"
	authServicePort "app/port/auth"
	passwordServicePort "app/port/passwordService"
	tenantServicePort "app/port/tenant"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
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
		//GetSingleService IGetSingleTenantAgent
		// GetSingleUserService authServicePort.IGetSingleUserService
		GetSingleService  tenantServicePort.IGetSingleTenant
		CreateUserService authServicePort.ICreateUser
		TenantAgentRepo   repository.ITenantAgent
		PasswordService   passwordServicePort.IPassword
		MongoClient       *mongo.Client
	}
)

func (this CreateTenantAgentService) Serve(inputUser *model.User, tenantUUID uuid.UUID, ctx context.Context) (*model.User, *model.TenantAgent, error) {

	if ctx == nil {

		ctx = context.Background()
	}

	if !libContext.IsNoRead(ctx) {

		existTenant, err := this.GetSingleService.CheckExist(tenantUUID, ctx)

		if err != nil {

			return nil, nil, err
		}

		if !existTenant {

			return nil, nil, errors.Join(
				common.ERR_NOT_FOUND, fmt.Errorf("tenant not found"),
			)
		}
	}

	inputUser.TenantUUID = &tenantUUID
	inputUser.UUID = libCommon.PointerPrimitive(uuid.New())

	newTenantAgent := &model.TenantAgent{
		UUID:       libCommon.PointerPrimitive(uuid.New()),
		TenantUUID: libCommon.PointerPrimitive(tenantUUID),
		UserUUID:   inputUser.UUID,
		CreatedBy:  inputUser.CreatedBy,
	}

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
		options.Transaction().
			SetWriteConcern(writeconcern.Majority()).
			SetReadConcern(readconcern.Snapshot()),
	)

	if err != nil {

		return nil, nil, err
	}
	//return this.GetSingleTenantService.Serve(newAgentModel.UUID.String())
	return inputUser, newTenantAgent, nil
}
