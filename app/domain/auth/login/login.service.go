package loginDomain

import (
	"app/internal/common"
	"app/model"
	accessTokenServicePort "app/port/accessToken"
	authServicePort "app/port/auth"
	authSignatureTokenPort "app/port/authSignatureToken"
	passwordServicePort "app/port/passwordService"
	refreshTokenServicePort "app/port/refreshToken"
	"app/repository"
	"context"
	"errors"
)

var (
	ERR_LOGIN_USER_NOT_FOUND = errors.New("loginService error: wrong username or password")
)

type (
	//ILogIn = authServiceAdapter.ILogIn

	LogInService struct {
		PasswordService         passwordServicePort.IPassword
		UserRepo                repository.IUser
		AccessTokenManipulator  accessTokenServicePort.IAccessTokenManipulator
		GetSingleUser           authServicePort.IGetSingleUser
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator

		AuthSignatureTokenProvider authSignatureTokenPort.IAuthSignatureProvider
	}
)

func (this *LogInService) Serve(
	username string, password string, ctx context.Context,
) (at string, rt string, err error) {

	existingUser, err := this.GetSingleUser.SearchByUsername(username, ctx)

	switch {
	case err != nil:
		return
	case existingUser == nil:
		err = errors.Join(
			common.ERR_NOT_FOUND,
			errors.New("user not found"),
		) // ERR_LOGIN_USER_NOT_FOUND
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

	return this.AuthSignatureTokenProvider.GenerateStrings(*existingUser.UUID, ctx)
}
