package tenantAgentService

import (
	"app/domain/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	IGetSingleTenantAgent interface {
		Serve(uuid string) (*model.TenantAgent, error)
		SearchByUsername(username string) (*model.TenantAgent, error)
		CheckUsernameExistence(username string) (bool, error)
	}

	GetSingleTenantAgentService struct {
		TenantAgentRepo repository.ITenantAgent
	}
)

func (this *GetSingleTenantAgentService) Serve(uuid_str string) (*model.TenantAgent, error) {

	userUUID, err := uuid.Parse(uuid_str)

	if err != nil {

		return nil, err
	}

	return this.TenantAgentRepo.FindOneByUUID(userUUID, context.TODO())
}

func (this *GetSingleTenantAgentService) SearchByUsername(username string) (*model.TenantAgent, error) {

	ret, err := this.TenantAgentRepo.Find(
		bson.D{
			{"username", username},
		},
		context.TODO(),
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetSingleTenantAgentService) CheckUsernameExistence(username string) (bool, error) {

	ret, err := this.TenantAgentRepo.Find(
		bson.D{
			{"username", username},
		},
		context.TODO(),
	)

	if err != nil {

		return false, err
	}

	if ret == nil {

		return false, nil
	}

	return true, nil
}
