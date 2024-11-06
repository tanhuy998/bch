package rotateSignaturesDomain

import (
	"app/internal/bootstrap"
	"app/internal/cacheList"
	"app/internal/common"
	"app/internal/generalToken"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	authServicePort "app/port/auth"
	authSignaturesServicePort "app/port/authSignatures"
	refreshTokenServicePort "app/port/refreshToken"
	refreshTokenClientPort "app/port/refreshTokenClient"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"context"
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ERR_REFRESH_NO_CONTEXT = errors.New("refresh login usecase: no context")
)

type (
	RefreshTokenBlackList = cacheList.CacheListManipulator[string, bootstrap.RefreshTokenBlackListCacheValue]
	GeneralTokenWhiteList = cacheList.CacheListManipulator[generalToken.GeneralTokenID, bootstrap.GeneralTokenWhiteListCacheValue]
)

type (
	RotateSignaturesUseCase struct {
		usecasePort.MongoUserSessionCacheUseCase[responsePresenter.RefreshLoginResponse]
		usecasePort.UseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
		GetSingleUserService   authServicePort.IGetSingleUser
		RefreshTokenIDProvider refreshTokenIdServicePort.IRefreshTokenIDProvider
		RefreshLoginService    authSignaturesServicePort.IRotateSignatures
		AccessTokenManipulator accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenClient     refreshTokenClientPort.IRefreshTokenClient
	}
)

func (this *RotateSignaturesUseCase) Execute(
	input *requestPresenter.RefreshLoginRequest,
) (output *responsePresenter.RefreshLoginResponse, err error) {

	reqCtx := input.GetContext()

	if reqCtx == nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(fmt.Errorf("refreshLoginUseCase: nil context given")),
		)
	}

	if v, ok := reqCtx.(iris.Context); ok {

		b, _ := v.GetBody()

		fmt.Println(string(b))
	}

	oldRefreshToken, err := this.RefreshTokenClient.Read(reqCtx)

	if err != nil {

		return nil, err
	}

	err = this.checkUserSession(input, oldRefreshToken)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	accessToken, err := this.AccessTokenManipulator.Read(input.Data.AccessToken)

	if err != nil {

		return nil, err
	}

	newAccessToken, newRefreshToken, err := this.RefreshLoginService.Serve(accessToken, oldRefreshToken, reqCtx)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	defer func() {

		if err != nil {

			return
		}

		err = this.RefreshTokenClient.Write(reqCtx, newRefreshToken)

		if err != nil {

			output = nil
			return
		}

		err = this.revokeRefreshToken(oldRefreshToken, input.GetContext())

		if err != nil {

			output = nil
			return
		}
	}()

	output = this.GenerateOutput()

	at, err := this.AccessTokenManipulator.SignString(newAccessToken)

	if err != nil {

		return nil, err
	}

	user, _ := this.GetSingleUserService.Serve(
		newAccessToken.GetUserUUID(), input.GetContext(),
	)

	output.Message = "success"
	output.Data = &responsePresenter.RefreshLoginData{
		AccessToken: at,
		User:        user,
	}

	return output, nil
}

func (this *RotateSignaturesUseCase) checkUserSession(input *requestPresenter.RefreshLoginRequest, refreshToken refreshTokenServicePort.IRefreshToken) error {

	if refreshToken == nil || refreshToken.Expired() {

		return errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("missing refresh token"),
		)
	}

	generalTokenID, _, err := this.RefreshTokenIDProvider.Extract(refreshToken.GetTokenID())

	if err != nil {

		return err
	}

	existsUserSession, err := this.GeneralTokenWhiteList.Has(generalTokenID, input.GetContext())

	if err != nil {

		return err
	}

	if existsUserSession {

		return nil
	}

	dbUserSession, err := this.UserSessionRepo.Find(
		bson.D{
			{"userUUID", generalTokenID.GetUserUUID()},
			{"tenantUUID", refreshToken.GetTenantUUID()},
			{"sessionID", generalTokenID},
		},
		input.GetContext(),
	)

	if err != nil {

		return err
	}

	if dbUserSession == nil {

		return common.ERR_UNAUTHORIZED
	}

	return nil
}

func (this *RotateSignaturesUseCase) revokeRefreshToken(
	refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context,
) error {

	if refreshToken == nil {

		return errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("missing refresh token"),
		)
	}
	fmt.Println(this.RefreshTokenBlackList, refreshToken.GetExpireTime())
	err := this.RefreshTokenBlackList.SetWithExpire(
		refreshToken.GetTokenID(), struct{}{}, *refreshToken.GetExpireTime(), ctx,
	)

	if err != nil {

		return err
	}

	return nil
}
