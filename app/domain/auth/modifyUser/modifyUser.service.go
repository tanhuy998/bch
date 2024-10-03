package modifyUserDomain

import (
	"app/model"
	authServicePort "app/port/auth"
	passwordServicePort "app/port/passwordService"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_MODIFY_USER_NOT_FOUND = errors.New("ModifyUser: user not found")
)

type (
	// IModifyUser interface {
	// 	Serve(userUUID string, data *model.User) error
	// }

	ModifyUserService struct {
		UserRepo      repository.IUser
		GetSingleUser authServicePort.IGetSingleUser
		Password      passwordServicePort.IPassword
	}
)

func (this *ModifyUserService) Serve(userUUID uuid.UUID, dataModel *model.User, ctx context.Context) error {

	// userUUID, err := uuid.Parse(userUUID_str)

	// if err != nil {

	// 	return err
	// }

	user, err := this.GetSingleUser.Serve(userUUID, ctx)

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
		userUUID, dataModel, ctx,
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
