package tenantService

import (
	tenantAgentServiceAdapter "app/adapter/tenantAgentService"
	"app/domain/model"
	libCommon "app/lib/common"
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
		GetSingleTenantAgentService tenantAgentServiceAdapter.IGetSingleTenantAgentServiceAdapter
		GetSingleTenantService      IGetSingleTenant
		TenantRepo                  repository.ITenant
		TenantAgentRepo             repository.ITenantAgent
		UserRepo                    repository.IUser
		MongoClient                 *mongo.Client
	}
)

func (this *CreateTenantService) Serve(
	name string, description string, tenantAgentUUID string,
) (*model.Tenant, error) {

	err := this.validateTenant(name)

	if err != nil {

		return nil, err
	}

	tenantAgent, err := this.fetchTenantAgent(tenantAgentUUID)

	if err != nil {

		return nil, err
	}

	session, err := this.MongoClient.StartSession()

	if err != nil {

		return nil, err
	}

	defer session.EndSession(context.TODO())

	result, err := session.WithTransaction(
		context.TODO(),
		func(sessionCtx mongo.SessionContext) (interface{}, error) {

			newTenant := &model.Tenant{
				UUID:        libCommon.PointerPrimitive(uuid.New()),
				Name:        name,
				Description: description,
			}

			err = this.TenantRepo.Create(newTenant, sessionCtx)

			if err != nil {

				return nil, err
			}

			err = this.update(tenantAgent, *newTenant.UUID, sessionCtx)

			if err != nil {

				sessionCtx.AbortTransaction(context.TODO())

				return nil, err
			}

			return newTenant, nil
		},
	)

	if err != nil {

		return nil, err
	}

	if val, ok := result.(*model.Tenant); ok {

		return val, nil
	}

	return nil, ERR_INTERNAL
}

func (this *CreateTenantService) validateTenant(name string) error {

	tenant, err := this.TenantRepo.Find(
		bson.D{
			{"name", name},
		},
		context.TODO(),
	)

	if err != nil {

		return err
	}

	if tenant != nil {

		return ERR_TENANT_EXISTS
	}

	return nil
}

func (this *CreateTenantService) fetchTenantAgent(tenantAgentUUID_str string) (*model.TenantAgent, error) {

	tenantAgent, err := this.GetSingleTenantAgentService.Serve(tenantAgentUUID_str)

	if err != nil {

		return nil, err
	}

	if tenantAgent == nil {

		return nil, ERR_NO_TENANT_AGENT
	}

	if tenantAgent.Deactivated {

		return nil, ERR_TENANT_AGENT_DEACTIVATED
	}

	if tenantAgent.TenantUUID != nil &&
		*tenantAgent.TenantUUID != uuid.Nil {

		return nil, ERR_INVALID_TENANT_AGENT
	}

	return tenantAgent, nil
}

func (this *CreateTenantService) update(tenantAgentModel *model.TenantAgent, tenantUUID uuid.UUID, ctx context.Context) error {

	g_err := ERR_INTERNAL

	//tenantAgentModel.TenantUUID = &tenantUUID

	// update tenantAgent repo
	for attempt := 0; attempt < RETRY_ATTEMPT; attempt++ {

		err := this.TenantAgentRepo.UpdateOneByUUID(
			*tenantAgentModel.UUID,
			&model.TenantAgent{
				TenantUUID: &tenantUUID,
			},
			ctx,
		)

		if err != nil {

			continue

		} else {

			g_err = nil
			break
		}
	}

	// if g_err != nil {

	// 	return g_err
	// }

	// // update user repo
	// for attempt := 0; attempt < RETRY_ATTEMPT; attempt++ {

	// 	err := this.UserRepo.UpdateOneByUUID(
	// 		tenantAgentModel.UserUUID,
	// 		&model.User{
	// 			TenantUUID: tenantUUID,
	// 		},

	// 		ctx,
	// 	)

	// 	if err != nil {

	// 		continue

	// 	} else {

	// 		g_err = nil
	// 		break
	// 	}
	// }

	return g_err
}
