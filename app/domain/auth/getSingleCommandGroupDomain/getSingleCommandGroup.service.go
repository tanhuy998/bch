package getSingleCommandGroupDomain

import (
	"app/model"
	"app/repository"
	repositoryAPI "app/repository/api"
	"context"

	"github.com/google/uuid"
)

type (
	GetSingleCommandGroupService struct {
		CommandGroupRepo repository.ICommandGroup
	}
)

func (this *GetSingleCommandGroupService) Serve(uuid uuid.UUID, ctx context.Context) (*model.CommandGroup, error) {

	return this.CommandGroupRepo.FindOneByUUID(uuid, ctx)
}

func (this *GetSingleCommandGroupService) SearchByName(groupName string, ctx context.Context) (*model.CommandGroup, error) {

	// ret, err := this.CommandGroupRepo.Find(
	// 	bson.D{
	// 		{"name", groupName},
	// 	},
	// 	ctx,
	// )

	ret, err := this.CommandGroupRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {
			filter.Field("name").Equal(groupName)
		},
	).FindOne(ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetSingleCommandGroupService) CheckCommandGroupExistence(groupName string, ctx context.Context) (bool, error) {

	// res, err := this.CommandGroupRepo.Find(
	// 	bson.D{
	// 		{"name", groupName},
	// 	},
	// 	ctx,
	// )

	res, err := this.CommandGroupRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("name").Equal(groupName)
		},
	).FindOne(ctx)

	if err != nil {

		return false, err
	}

	if res == nil {

		return false, nil
	}

	return true, nil
}
