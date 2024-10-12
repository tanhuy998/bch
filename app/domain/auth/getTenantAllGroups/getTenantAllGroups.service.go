package getTenantAllGroupsDomain

import (
	"app/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	GetTenantAllGroupService struct {
		CommandGroupRepo repository.ICommandGroup
	}
)

func (this *GetTenantAllGroupService) Serve(tenantUUID uuid.UUID, ctx context.Context) ([]*model.CommandGroup, error) {

	return repository.Aggregate[model.CommandGroup](
		this.CommandGroupRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"tenantUUID", tenantUUID},
					},
				},
			},
		},
		ctx,
	)
}
