package getTenantAllGroupsDomain

import (
	"app/model"
	"app/repository"
	repositoryAPI "app/repository/api"
	"context"

	"github.com/google/uuid"
)

type (
	GetTenantAllGroupService struct {
		CommandGroupRepo repository.ICommandGroup
	}
)

func (this *GetTenantAllGroupService) Serve(tenantUUID uuid.UUID, ctx context.Context) ([]*model.CommandGroup, error) {

	// return repository.Aggregate[model.CommandGroup](
	// 	this.CommandGroupRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{"tenantUUID", tenantUUID},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	ctx,
	// )

	return this.CommandGroupRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {
			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	).Find(ctx)
}
