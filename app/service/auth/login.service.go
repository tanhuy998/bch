package authService

import (
	accessTokenServicePort "app/adapter/accessToken"
	authServiceAdapter "app/adapter/auth"
	passwordServiceAdapter "app/adapter/passwordService"
	refreshTokenServicePort "app/adapter/refreshToken"
	"app/domain/model"
	"app/repository"
	"context"
	"errors"
	"fmt"
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
	}
)

func (this *LogInService) Serve(
	username string, password string, ctx context.Context,
) (at string, rt string, err error) {

	existingUser, err := this.GetSingleUser.SearchByUsername(username, ctx)

	if err != nil {

		return
	}

	if existingUser == nil {

		err = ERR_LOGIN_USER_NOT_FOUND
		return
	}

	inputModel := &model.User{
		Username: username,
		PassWord: password,
	}

	err = this.PasswordService.Resolve(inputModel)

	if err != nil {

		return
	}

	if this.PasswordService.Compare(inputModel, existingUser.GetSecret()) != nil {

		err = ERR_LOGIN_USER_NOT_FOUND
		return
	}

	accessToken, err := this.AccessTokenManipulator.GenerateByUserUUID(existingUser.UUID, ctx)

	if err != nil {

		return
	}

	fmt.Println(3, accessToken.GetUserUUID())
	refreshToken, err := this.RefreshTokenManipulator.Generate(accessToken.GetUserUUID(), ctx)

	if err != nil {

		return
	}

	fmt.Println(4)
	at, err = this.AccessTokenManipulator.SignString(accessToken)

	if err != nil {

		return
	}

	rt, err = this.RefreshTokenManipulator.SignString(refreshToken)
	fmt.Println(5)
	if err != nil {

		return
	}

	return
}

// func (this *LogInService) generateAccessToken(authData *valueObject.AuthData, ctx context.Context) (string, error) {

// 	this.AccessTokenManipulator.Generate(*authData.UserUUID, ctx)
// }

// func (this *LogInService) generateRefreshToken(authData *valueObject.AuthData) string {

// }
