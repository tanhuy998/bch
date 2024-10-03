package getSingleCommandGroupDomain

import (
	"app/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	// IGetSingleCommandGroup interface {
	// 	Serve(uuid uuid.UUID, ctx context.Context) (*model.CommandGroup, error)
	// 	SearchByName(groupName string, ctx context.Context) (*model.CommandGroup, error)
	// 	CheckCommandGroupExistence(groupName string, ctx context.Context) (bool, error)
	// }

	GetSingleCommandGroupService struct {
		CommandGroupRepo repository.ICommandGroup
	}
)

func (this *GetSingleCommandGroupService) Serve(uuid uuid.UUID, ctx context.Context) (*model.CommandGroup, error) {

	return this.CommandGroupRepo.FindOneByUUID(uuid, ctx)
}

func (this *GetSingleCommandGroupService) SearchByName(groupName string, ctx context.Context) (*model.CommandGroup, error) {

	ret, err := this.CommandGroupRepo.Find(
		bson.D{
			{"name", groupName},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetSingleCommandGroupService) CheckCommandGroupExistence(groupName string, ctx context.Context) (bool, error) {

	res, err := this.CommandGroupRepo.Find(
		bson.D{
			{"name", groupName},
		},
		ctx,
	)

	if err != nil {

		return false, err
	}

	if res == nil {

		return false, nil
	}

	return true, nil
}
