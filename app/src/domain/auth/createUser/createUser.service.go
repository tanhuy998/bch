package createUserDomain

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	authServicePort "app/src/port/auth"
	passwordServicePort "app/src/port/passwordService"
	"app/src/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_USER_NAME_EXISTS = errors.New("username already exists.")
)

type (
	// ICreateUser interface {
	// 	Serve(username string, password string, name string, ctx context.Context) (*model.User, error)
	// 	CreateByModel(dataModel *model.User, ctx context.Context) (*model.User, error)
	// }

	CreateUserService struct {
		UserRepo        repository.IUser
		GetSingleUser   authServicePort.IGetSingleUser
		PasswordAdapter passwordServicePort.IPassword
	}
)

func (this *CreateUserService) CreateByModel(model *model.User, ctx context.Context) (*model.User, error) {

	usernameExist, err := this.GetSingleUser.CheckUsernameExistence(model.Username, ctx)

	if err != nil {

		return nil, err
	}

	if usernameExist {

		return nil, ERR_USER_NAME_EXISTS
	}

	err = this.PasswordAdapter.Resolve(model)

	if err != nil {

		return nil, err
	}

	model.UUID = libCommon.PointerPrimitive(uuid.New())

	err = this.UserRepo.Create(model, ctx)

	if err != nil {

		return nil, err
	}

	return model, nil
}

func (this *CreateUserService) Serve(
	username string,
	password string,
	name string,
	ctx context.Context,
) (*model.User, error) {

	model := &model.User{
		//UUID:     uuid.New(),
		Username: username,
		PassWord: password,
		Name:     name,
	}

	return this.CreateByModel(model, ctx)
}
