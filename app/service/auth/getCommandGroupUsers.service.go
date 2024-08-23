package authService

import (
	"app/domain/model"
	"app/repository"
	"context"
	"errors"

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

func (this *GetCommandGroupUsersService) Serve(groupUUID_str string) ([]*model.User, error) {

	groupUUID, err := uuid.Parse(groupUUID_str)

	if err != nil {

		return nil, ERR_INVALID_INPUT_GROUP
	}

	group, err := this.CommandGroupRepo.FindOneByUUID(groupUUID, context.TODO())

	if err != nil {

		return nil, err
	}

	if group == nil {

		return nil, ERR_NO_GROUP
	}

	ret, err := repository.Aggregate[model.User](
		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"commandGroupUUID", groupUUID},
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
		context.TODO(),
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}
