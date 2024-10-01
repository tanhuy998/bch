package authService

import (
	"app/src/model"
	"app/src/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	ICheckUserInCommandGroup interface {
		Serve(groupUUID, userUUID string) (bool, error)
		Detail(groupUUID uuid.UUID, userUUID uuid.UUID) (*model.CommandGroupUser, error)
	}

	CheckUserInCommandGroupService struct {
		CommandGroupUserRepo repository.ICommandGroupUser
	}
)

func (this *CheckUserInCommandGroupService) Serve(groupUUID_str, userUUID_str string) (bool, error) {

	groupUUID, err := uuid.Parse(groupUUID_str)

	if err != nil {

		return false, nil
	}

	userUUID, err := uuid.Parse(userUUID_str)

	if err != nil {

		return false, nil
	}

	res, err := this.CommandGroupUserRepo.Find(
		bson.D{
			{"commandGroupUUID", groupUUID},
			{"userUUID", userUUID},
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

func (this *CheckUserInCommandGroupService) Detail(groupUUID uuid.UUID, userUUID uuid.UUID) (*model.CommandGroupUser, error) {

	ret, err := this.CommandGroupUserRepo.Find(
		bson.D{
			{"commandGroupUUID", groupUUID},
			{"userUUID", userUUID},
		},
		context.TODO(),
	)

	if err != nil {

		return ret, err
	}

	return ret, nil
}
