package grantUserAsTenantAgentDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
	tenantServicePort "app/port/tenant"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	SERVICE_NAME = "GrantUserAsTenantService"
)

type (
	GrantUserAsTenantAgentService struct {
		GetSingleTenantService tenantServicePort.IGetSingleTenant
		GetSingleUserService   authServicePort.IGetSingleUser
		TenantAgentRepo        repository.ITenantAgent
		MongoClient            *mongo.Client
	}
)

func (this *GrantUserAsTenantAgentService) Serve(
	userUUID uuid.UUID, tenantUUID uuid.UUID, newTenantAgent *model.TenantAgent, ctx context.Context,
) (*model.TenantAgent, error) {

	if ctx == nil {

		ctx = context.Background()
	}

	existTenant, err := this.GetSingleTenantService.CheckExist(tenantUUID, ctx)

	if err != nil {

		return nil, err
	}

	if !existTenant {

		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("%s error: tenant not found", SERVICE_NAME))
	}

	existUser, err := this.GetSingleTenantService.CheckExist(userUUID, ctx)

	if err != nil {

		return nil, err
	}

	if !existUser {

		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("%s error: user not found", SERVICE_NAME))
	}

	userAlreadyTenantAgent, err := this.TenantAgentRepo.Find(
		bson.D{
			{"userUUID", userUUID},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	if userAlreadyTenantAgent != nil {

		return nil, errors.Join(common.ERR_BAD_REQUEST, fmt.Errorf("%s error: user are currently a tenant agent", SERVICE_NAME))
	}

	newTenantAgent.TenantUUID = libCommon.PointerPrimitive(tenantUUID)

	err = this.TenantAgentRepo.Create(newTenantAgent, ctx)

	if err != nil {

		return nil, err
	}

	return newTenantAgent, nil
}
