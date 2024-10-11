package getCommandGroupUsersDomain

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

var (
	ERR_INVALID_INPUT_GROUP = errors.New("GetGroupUsers: invalid group")
	ERR_NO_GROUP            = errors.New("GetGroupUser: command group not found")
)

type (
	IGetCommandGroupUsers interface {
		Serve(groupUUID string) ([]*model.User, error)
	}

	GetCommandGroupUsersService struct {
		CommandGroupUserRepo repository.ICommandGroupUser
		CommandGroupRepo     repository.ICommandGroup
	}
)

func (this *GetCommandGroupUsersService) Serve(
	tenantUUID uuid.UUID, groupUUID uuid.UUID, ctx context.Context,
) ([]*model.User, error) {

	switch existingGroup, err := this.CommandGroupRepo.FindOneByUUID(groupUUID, ctx); {
	case err != nil:
		return nil, err
	case existingGroup == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("group not found"))
	case *existingGroup.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("group not in tenant"))
	}

	ret, err := repository.Aggregate[model.User](
		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"tenantUUID", tenantUUID},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "users"},
						{"localField", "userUUID"},
						{"foreignField", "uuid"},
						{"as", "users"},
					},
				},
			},
			bson.D{{"$unwind", bson.D{{"path", "$users"}}}},
			bson.D{{"$replaceRoot", bson.D{{"newRoot", "$$ROOT.users"}}}},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

// func (this *GetCommandGroupUsersService) SearchAndRetrieveByModel(dataModel *model.CommandGroup, ctx context.Context) ([]*model.User, error) {

// 	groupUUID := dataModel.UUID
// 	group, err := this.CommandGroupRepo.FindOneByUUID(*groupUUID, context.TODO())

// 	if err != nil {

// 		return nil, err
// 	}

// 	if group == nil {

// 		return nil, ERR_NO_GROUP
// 	}

// 	ret, err := repository.Aggregate[model.User](
// 		this.CommandGroupUserRepo.GetCollection(),
// 		mongo.Pipeline{
// 			bson.D{
// 				{
// 					"$match", dataModel,
// 				},
// 			},
// 			bson.D{
// 				{"$lookup",
// 					bson.D{
// 						{"from", "users"},
// 						{"localField", "userUUID"},
// 						{"foreignField", "uuid"},
// 						{"as", "users"},
// 					},
// 				},
// 			},
// 			bson.D{{"$unwind", bson.D{{"path", "$users"}}}},
// 			bson.D{{"$replaceRoot", bson.D{{"newRoot", "$$ROOT.users"}}}},
// 		},
// 		context.TODO(),
// 	)

// 	if err != nil {

// 		return nil, err
// 	}

// 	return ret, nil
// }
