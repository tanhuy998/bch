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
		Serve(uuid string, ctx context.Context) (*model.User, error)
		SearchByUsername(username string, ctx context.Context) (*model.User, error)
		CheckUsernameExistence(username string, ctx context.Context) (bool, error)
	}

	GetSingleUser struct {
		UserRepo repository.IUser
	}
)

func (this *GetSingleUser) Serve(uuid_str string, ctx context.Context) (*model.User, error) {

	userUUID, err := uuid.Parse(uuid_str)

	if err != nil {

		return nil, err
	}

	return this.UserRepo.FindOneByUUID(userUUID, ctx)
}

func (this *GetSingleUser) SearchByUsername(username string, ctx context.Context) (*model.User, error) {

	ret, err := this.UserRepo.Find(
		bson.D{
			{"username", username},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetSingleUser) CheckUsernameExistence(username string, ctx context.Context) (bool, error) {

	ret, err := this.UserRepo.Find(
		bson.D{
			{"username", username},
		},
		ctx,
	)

	if err != nil {

		return false, err
	}

	if ret == nil {

		return false, nil
	}

	return true, nil
}
