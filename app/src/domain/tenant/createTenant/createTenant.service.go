package createTenantDomain

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	tenantServicePort "app/src/port/tenant"
	"app/src/repository"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	RETRY_ATTEMPT = 3
)

var (
	ERR_NO_TENANT_AGENT          = errors.New("createTenantService: tenant agent not found")
	ERR_TENANT_EXISTS            = errors.New("createTenantService: tenant already exists")
	ERR_INVALID_TENANT_AGENT     = errors.New("createTenantService: invlaid tenant agent")
	ERR_TENANT_AGENT_DEACTIVATED = errors.New("createTenantService: tenant agent deactivated")
	ERR_INTERNAL                 = errors.New("createTenantService: internal error")
)

type (
	ICreateTenant interface {
		Serve(name string, description string, tenantAgentUUID string) (*model.Tenant, error)
	}

	CreateTenantService struct {
		CreateTenantAgentService tenantServicePort.ICreateTenantAgent
		//GetSingleTenantService   tenantServicePort.IGetSingleTenant
		TenantRepo      repository.ITenant
		TenantAgentRepo repository.ITenantAgent
		UserRepo        repository.IUser
		MongoClient     *mongo.Client
	}
)

func (this *CreateTenantService) Serve(
	inputTenant *model.Tenant, inputUser *model.User, ctx context.Context,
) (t *model.Tenant, u *model.User, err error) {

	err = this.validateTenant(inputTenant.Name, ctx)

	if err != nil {

		return
	}

	session, err := this.MongoClient.StartSession()

	if err != nil {

		return
	}

	defer session.EndSession(context.TODO())

	_, err = session.WithTransaction(
		ctx,
		func(sessionCtx mongo.SessionContext) (any, error) {

			inputTenant.UUID = libCommon.PointerPrimitive(uuid.New())

			err := this.TenantRepo.Create(inputTenant, sessionCtx)

			if err != nil {

				return nil, err
			}

			inputUser, _, err = this.CreateTenantAgentService.Serve(inputUser, *inputTenant.UUID, sessionCtx)

			if err != nil {

				return nil, err
			}

			return nil, nil
		},
	)

	if err != nil {

		return
	}

	return inputTenant, inputUser, nil
}

func (this *CreateTenantService) validateTenant(name string, ctx context.Context) error {

	tenant, err := this.TenantRepo.Find(
		bson.D{
			{"name", name},
		},
		ctx,
	)

	if err != nil {

		return err
	}

	if tenant != nil {

		return ERR_TENANT_EXISTS
	}

	return nil
}
