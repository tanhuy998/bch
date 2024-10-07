package createTenantDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	tenantServicePort "app/port/tenant"
	"app/repository"
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

	if ctx != nil {

		ctx = context.Background()
	}

	err = this.validateTenant(inputTenant.Name, ctx)

	if err != nil {

		return
	}

	session, err := this.MongoClient.StartSession()

	if err != nil {

		err = errors.Join(common.ERR_INTERNAL, err)
		return
	}

	defer session.EndSession(ctx)

	newAgentModel := &model.TenantAgent{
		UUID:       libCommon.PointerPrimitive(uuid.New()),
		UserUUID:   libCommon.PointerPrimitive(uuid.UUID(*inputUser.UUID)),
		TenantUUID: inputUser.TenantUUID,
	}

	_, err = session.WithTransaction(
		ctx,
		func(sessionCtx mongo.SessionContext) (any, error) {

			inputTenant.UUID = libCommon.PointerPrimitive(uuid.New())

			err := this.TenantRepo.Create(inputTenant, sessionCtx)

			if err != nil {

				return nil, err
			}

			if inputUser == nil {

				return nil, nil
			}

			inputUser, _, err = this.CreateTenantAgentService.Serve(inputUser, newAgentModel, *inputTenant.UUID, sessionCtx)

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

		return errors.Join(common.ERR_INTERNAL, err)
	}

	if tenant != nil {

		return ERR_TENANT_EXISTS
	}

	return nil
}
