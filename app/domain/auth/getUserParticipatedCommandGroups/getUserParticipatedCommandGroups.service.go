package getUserParticipatedCommandGroupsDomain

import (
	"app/internal/common"
	"app/model"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	GetUserParticipatedCommandGroupService struct {
		UserRepo             repository.IUser
		CommandGroupUserRepo repository.ICommandGroupUser
	}
)

func (this *GetUserParticipatedCommandGroupService) Serve(tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context) ([]*model.CommandGroup, error) {

	switch existingUser, err := this.UserRepo.FindOneByUUID(userUUID, ctx); {
	case err != nil:
		return nil, err
	case existingUser == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("user not found"))
	case *existingUser.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("user not in tenant"))
	}

	ret, err := repository.Aggregate[model.CommandGroup](
		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"userUUID", userUUID},
						{"tenantUUID", tenantUUID},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroups"},
						{"localField", "commandGroupUUID"},
						{"foreignField", "uuid"},
						{"as", "commandGroups"},
					},
				},
			},
			bson.D{{"$unwind", "$commandGroups"}},
			bson.D{{"$replaceWith", "$commandGroups"}},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetUserParticipatedCommandGroupService) SearchAndRetrieveByModel(searchModel *model.User, ctx context.Context) ([]*model.CommandGroup, error) {

	return nil, nil
}
