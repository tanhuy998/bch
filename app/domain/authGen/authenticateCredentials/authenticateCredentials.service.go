package authenticateCredentialsDomain

import (
	"app/internal/common"
	"app/model"
	authServicePort "app/port/auth"
	generalTokenServicePort "app/port/generalToken"

	passwordServicePort "app/port/passwordService"
	"app/repository"
	"context"
	"errors"
)

var (
	ERR_LOGIN_USER_NOT_FOUND = errors.New("AuthenticateCredentialsService error: wrong username or password")
)

type (
	AuthenticateCredentialsService struct {
		PasswordService     passwordServicePort.IPassword
		UserRepo            repository.IUser
		GetSingleUser       authServicePort.IGetSingleUser
		GeneralTokenService generalTokenServicePort.IGeneralTokenManipulator
	}
)

func (this *AuthenticateCredentialsService) Serve(
	username string, password string, ctx context.Context,
) (gt generalTokenServicePort.IGeneralToken, err error) {

	existingUser, err := this.GetSingleUser.SearchByUsername(username, ctx)

	switch {
	case err != nil:
		return
	case existingUser == nil:
		err = errors.Join(
			common.ERR_NOT_FOUND,
			errors.New("user not found"),
		)
		return
	}

	inputModel := &model.User{
		Username: username,
		PassWord: password,
	}

	err = this.PasswordService.Resolve(inputModel)

	switch {
	case err != nil:
		return
	case this.PasswordService.Compare(inputModel, existingUser.GetSecret()) != nil:
		err = errors.New("invalid user credentials")
		return
	}

	generalToken, err := this.GeneralTokenService.Generate(*existingUser.UUID, ctx)

	if err != nil {

		return
	}

	return generalToken, nil
}
