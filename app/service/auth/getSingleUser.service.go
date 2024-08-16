package authService

import (
	"app/domain/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	IGetSingleUser interface {
		Serve(uuid string) (*model.User, error)
		SearchByUsername(username string) (*model.User, error)
		CheckUsernameExistence(username string) (bool, error)
	}

	GetSingleUser struct {
		UserRepo repository.IUser
	}
)

func (this *GetSingleUser) Serve(uuid_str string) (*model.User, error) {

	userUUID, err := uuid.Parse(uuid_str)

	if err != nil {

		return nil, err
	}

	return this.UserRepo.FindOneByUUID(userUUID, context.TODO())
}

func (this *GetSingleUser) SearchByUsername(username string) (*model.User, error) {

	ret, err := this.UserRepo.Find(
		bson.D{
			{"username", username},
		},
		context.TODO(),
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetSingleUser) CheckUsernameExistence(username string) (bool, error) {

	ret, err := this.UserRepo.Find(
		bson.D{
			{"username", username},
		},
		context.TODO(),
	)

	if err != nil {

		return false, err
	}

	if ret == nil {

		return false, nil
	}

	return true, nil
}
