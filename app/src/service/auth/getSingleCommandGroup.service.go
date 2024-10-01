package authService

import (
	"app/src/model"
	"app/src/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	IGetSingleCommandGroup interface {
		Serve(uuid string) (*model.CommandGroup, error)
		SearchByName(groupName string) (*model.CommandGroup, error)
		CheckCommandGroupExistence(groupName string) (bool, error)
	}

	GetSingleCommandGroupService struct {
		CommandGroupRepo repository.ICommandGroup
	}
)

func (this *GetSingleCommandGroupService) Serve(uuid_str string) (*model.CommandGroup, error) {

	uuid, err := uuid.Parse(uuid_str)

	if err != nil {

		return nil, err
	}

	return this.CommandGroupRepo.FindOneByUUID(uuid, context.TODO())
}

func (this *GetSingleCommandGroupService) SearchByName(groupName string) (*model.CommandGroup, error) {

	ret, err := this.CommandGroupRepo.Find(
		bson.D{
			{"name", groupName},
		},
		context.TODO(),
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetSingleCommandGroupService) CheckCommandGroupExistence(groupName string) (bool, error) {

	res, err := this.CommandGroupRepo.Find(
		bson.D{
			{"name", groupName},
		},
		context.TODO(),
	)

	if err != nil {

		return false, err
	}

	if res == nil {

		return false, nil
	}

	return true, nil
}
