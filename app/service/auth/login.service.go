package authService

import (
	accessTokenServicePort "app/adapter/accessToken"
	authServiceAdapter "app/adapter/auth"
	authSignatureTokenPort "app/adapter/authSignatureToken"
	passwordServiceAdapter "app/adapter/passwordService"
	refreshTokenServicePort "app/adapter/refreshToken"
	"app/domain/model"
	"app/repository"
	"context"
	"errors"
)

var (
	ERR_LOGIN_USER_NOT_FOUND = errors.New("loginService error: wrong username or password")
)

type (
	ILogIn = authServiceAdapter.ILogIn

	LogInService struct {
		PasswordService         passwordServiceAdapter.IPassword
		UserRepo                repository.IUser
		AccessTokenManipulator  accessTokenServicePort.IAccessTokenManipulator
		GetSingleUser           authServiceAdapter.IGetSingleUserService
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
		err = ERR_LOGIN_USER_NOT_FOUND
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
		err = ERR_LOGIN_USER_NOT_FOUND
		return
	}

	return this.AuthSignatureTokenProvider.GenerateStrings(existingUser.UUID, ctx)
}
