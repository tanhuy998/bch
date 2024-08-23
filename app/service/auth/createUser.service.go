package authService

import (
	passwordAdapter "app/adapter/passwordService"
	"app/domain/model"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_USER_NAME_EXISTS = errors.New("username already exists.")
)

type (
	ICreateUser interface {
		Serve(username string, password string, name string) (*model.User, error)
	}

	CreateUserService struct {
		UserRepo        repository.IUser
		GetSingleUser   IGetSingleUser
		PasswordAdapter passwordAdapter.IPassword
	}
)

func (this *CreateUserService) Serve(
	username string,
	password string,
	name string,
) (*model.User, error) {

	usernameExist, err := this.GetSingleUser.CheckUsernameExistence(username)

	if err != nil {

		return nil, err
	}

	if usernameExist {

		return nil, ERR_USER_NAME_EXISTS
	}

	model := &model.User{
		UUID:     uuid.New(),
		Username: username,
		PassWord: password,
		Name:     name,
	}

	err = this.PasswordAdapter.Resolve(model)

	if err != nil {

		return nil, err
	}

	err = this.UserRepo.Create(model, context.TODO())

	if err != nil {

		return nil, err
	}

	return model, nil
}
