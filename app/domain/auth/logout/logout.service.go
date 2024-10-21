package logoutDomain

import (
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	authServicePort "app/port/auth"
	refreshTokenServicePort "app/port/refreshToken"
	"app/repository"
	"context"
	"fmt"
)

type (
	LogoutService struct {
		UserSessionRepo     repository.IUserSession
		RemoveDBUserSession authServicePort.IRemoveDBUserSession
	}
)

func (this *LogoutService) Serve(
	refreshToken refreshTokenServicePort.IRefreshToken, accessToken accessTokenServicePort.IAccessToken, ctx context.Context,
) (err error) {

	defer func() {

		if err != nil {

			return
		}

		err = this.RemoveDBUserSession.Serve(accessToken.GetUserUUID(), ctx)

		if err != nil {

			return
		}
	}()

	switch {
	case refreshToken == nil || accessToken == nil:
		return libError.NewInternal(
			fmt.Errorf("nil signature given"),
		)
	case refreshToken.GetTokenID() != accessToken.GetTokenID():
		return libError.NewInternal(
			fmt.Errorf("auth signatures not matches"),
		)
	case refreshToken.Expired():
		return nil
	}

	return nil
}
