package authService

import (
	passwordAdapter "app/adapter/passwordService"
	"app/domain/model"
	"app/repository"
	"context"
	"errors"
	"fmt"

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

	pw, err := this.PasswordAdapter.Generate(password)

	if err != nil {

		return nil, err
	}

	model := &model.User{
		UUID:     uuid.New(),
		Username: username,
		PassWord: pw,
	}

	fmt.Println(this.UserRepo.GetCollection().Name())
	err = this.UserRepo.Create(model, context.TODO())

	if err != nil {

		return nil, err
	}

	return model, nil
}
