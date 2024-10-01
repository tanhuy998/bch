package authService

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	"app/src/repository"
	authValueObject "app/src/valueObject/auth"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IIdentityManipulator interface {
		CreateUser(username string, password string) error
		AssignUserToCommandGroup(userUUID uuid.UUID, commandGroupUUID uuid.UUID) (*model.CommandGroupUser, error)
		GrantCommandGroupUserRole(commandGroupUserUUID uuid.UUID, RoleUUID uuid.UUID) error
		GetCommandeeCommandGroups(userUUID uuid.UUID) (*repository.PaginationPack[model.CommandGroup], error)
		GetGroupMembers(GroupUUID uuid.UUID, pivot primitive.ObjectID, limit int, isPrev bool) (*repository.PaginationPack[authValueObject.CommandGroupUserEntity], error)
	}

	AuthManipulator struct {
		UserRepo                 repository.IUser
		CommandGroupRepo         repository.ICommandGroup
		CommandGroupUserRepo     repository.ICommandGroupUser
		CommandGroupUserRoleRepo repository.ICommandGroupUserRole
		RoleRepo                 repository.IRole
	}
)

func (this *AuthManipulator) CreateUser(
	username string,
	password string,
	name string,
) (*model.User, error) {

	model := &model.User{
		UUID:     libCommon.PointerPrimitive(uuid.New()),
		Username: username,
		//PassWord: password,
	}

	err := this.UserRepo.Create(model, context.TODO())

	if err != nil {

		return nil, err
	}

	return model, nil
}

func (this *AuthManipulator) AssignUserToCommandGroup(
	userUUID uuid.UUID,
	commandGroupUUID uuid.UUID,
) (*model.CommandGroupUser, error) {

	user, err := this.UserRepo.FindOneByUUID(userUUID, context.TODO())

	if err != nil {

		return nil, err
	}

	if user == nil {

		return nil, nil //common.ERR_HTTP_NOT_FOUND
	}

	commandGroup, err := this.CommandGroupRepo.FindOneByUUID(commandGroupUUID, context.TODO())

	if err != nil {

		return nil, err
	}

	if commandGroup == nil {

		return nil, nil //common.ERR_HTTP_NOT_FOUND
	}

	model := &model.CommandGroupUser{
		UUID:             libCommon.PointerPrimitive(uuid.New()),
		UserUUID:         &userUUID,
		CommandGroupUUID: &commandGroupUUID,
	}

	err = this.CommandGroupUserRepo.Create(model, context.TODO())

	if err != nil {

		return nil, err
	}

	return model, nil
}

func (this *AuthManipulator) GrantCommandGroupUserRole(
	commandGroupUserUUID uuid.UUID,
	roleUUID uuid.UUID,
) error {

	role, err := this.RoleRepo.FindOneByUUID(roleUUID, context.TODO())

	if err != nil {

		return err
	}

	if role == nil {

		return errors.New("role not found") // common.ERR_HTTP_NOT_FOUND
	}

	commandGroupUser, err := this.CommandGroupUserRepo.FindOneByUUID(commandGroupUserUUID, context.TODO())

	if err != nil {

		return err
	}

	if commandGroupUser == nil {

		return errors.New("user has not been in command group yet") // common.ERR_HTTP_NOT_FOUND
	}

	model := &model.CommandGroupUserRole{
		UUID:                 libCommon.PointerPrimitive(uuid.New()),
		CommandGroupUserUUID: &commandGroupUserUUID,
		RoleUUID:             &roleUUID,
	}

	return this.CommandGroupUserRoleRepo.Create(model, context.TODO())
}

func (this *AuthManipulator) GetGroupMembers(
	groupUUID uuid.UUID,
	pivot primitive.ObjectID,
	limit int,
	isPrev bool,
) (*repository.PaginationPack[authValueObject.CommandGroupUserEntity], error) {

	repository.Aggregate[authValueObject.CommandGroupUserEntity](
		this.CommandGroupRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"$uuid", groupUUID},
					},
				},
			},
			bson.D{
				{
					"$lookup", bson.D{
						{"from", repository.COMMAND_GROUP_USER_COLLECTION_NAME},
						{"localField", "uuid"},
						{"foreignField", "commandGroupUserUUID"},
						{"as", "groupUsers"},
					},
				},
			},
			bson.D{
				{"$unwind", "groupUsers"},
			},
			bson.D{
				{
					"$set", bson.D{
						{"userUUID", "$groupUsers.userUUID"},
					},
				},
			},
			bson.D{
				{
					"$lookup", bson.D{
						{"from", repository.USER_COLLECTION_NAME},
						{"localField", "userUUID"},
						{"foreignField", "uuid"},
						{"as", "users"},
					},
				},
			},
			bson.D{
				{
					"$unwind", "users",
				},
			},
			// bson.D{
			// 	{
			// 		"$project", bson.D{
			// 			{""},
			// 		},
			// 	},
			// },
		},
		context.TODO(),
	)

	return nil, nil
}

func (this *AuthManipulator) GetCommandeeCommandGroups(
	userUUID uuid.UUID,
) (*repository.PaginationPack[model.CommandGroup], error) {

	return nil, nil
}
