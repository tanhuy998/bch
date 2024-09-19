package authService

import (
	passwordServiceAdapter "app/adapter/passwordService"
	"app/domain/model"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_MODIFY_USER_NOT_FOUND = errors.New("ModifyUser: user not found")
)

type (
	IModifyUser interface {
		Serve(userUUID string, data *model.User) error
	}

	ModifyUserService struct {
		UserRepo      repository.IUser
		GetSingleUser IGetSingleUser
		Password      passwordServiceAdapter.IPassword
	}
)

func (this *ModifyUserService) Serve(userUUID_str string, dataModel *model.User) error {

	userUUID, err := uuid.Parse(userUUID_str)

	if err != nil {

		return err
	}

	user, err := this.GetSingleUser.Serve(userUUID_str, context.TODO())

	if err != nil {

		return err
	}

	if user == nil {

		return ERR_MODIFY_USER_NOT_FOUND
	}

	dataModel.UUID = user.UUID
	dataModel.Username = user.Username

	err = this.resolvePassword(dataModel)

	if err != nil {

		return err
	}

	err = this.UserRepo.UpdateOneByUUID(
		userUUID, dataModel, context.TODO(),
	)

	if err != nil {

		return err
	}

	return nil
}

func (this *ModifyUserService) resolvePassword(dataModel *model.User) error {

	if dataModel.PassWord == "" {

		return nil
	}

	return this.Password.Resolve(dataModel)
}
