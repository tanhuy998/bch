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
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"
	"context"
	"errors"
	"fmt"

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
		//usecasePort.MongoUserSessionCacheUseCase[responsePresenter.RefreshLoginResponse]
		//usecasePort.UseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
		unitOfWork.GenericUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse]
		unitOfWork.MongoUserSessionCacheUseCase[responsePresenter.RefreshLoginResponse]
		unitOfWork.OperationLogger
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

	defer this.WrapResults(input, &output, &err)

	reqCtx := input.GetContext()

	if reqCtx == nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(fmt.Errorf("refreshLoginUseCase: nil context given")),
		)
	}

	oldRefreshToken, err := this.RefreshTokenClient.Read(reqCtx)

	this.PushTraceIfError(err, "read_refresh_token", "failed", input.GetContext())

	if err != nil {

		return nil, err
	}

	this.PushTrace("read_refresh_token", "", input.GetContext())

	err = this.checkUserSession(input, oldRefreshToken)

	if err != nil {

		// return nil, this.ErrorWithContext(
		// 	input, err,
		// )

		return
	}

	accessToken, err := this.AccessTokenManipulator.Read(input.Data.AccessToken)
	this.PushTraceIfError(err, "read_access_token", "failed", reqCtx)

	if err != nil {

		//return nil, err
		return
	}

	newAccessToken, newRefreshToken, err := this.RefreshLoginService.Serve(accessToken, oldRefreshToken, reqCtx)

	this.PushTraceIfError(err, "generat_new_signatures", "failed", reqCtx)
	if err != nil {

		//return nil, this.ErrorWithContext(input, err)
		return
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

		//return nil, err
		return
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

func (this *RotateSignaturesUseCase) checkUserSession(
	input *requestPresenter.RefreshLoginRequest, refreshToken refreshTokenServicePort.IRefreshToken,
) error {

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

	return this.resetCacheSession(input, refreshToken)
}

func (this *RotateSignaturesUseCase) resetCacheSession(
	input *requestPresenter.RefreshLoginRequest, refreshToken refreshTokenServicePort.IRefreshToken,
) error {

	generalTokenID, _, err := this.RefreshTokenIDProvider.Extract(refreshToken.GetTokenID())

	if err != nil {

		return err
	}

	exp := refreshToken.GetExpireTime()

	if exp == nil {

		_, err = this.GeneralTokenWhiteList.Set(generalTokenID, struct{}{}, input.GetContext())

	} else {

		err = this.GeneralTokenWhiteList.SetWithExpire(generalTokenID, struct{}{}, *exp, input.GetContext())
	}

	switch {
	case errors.Is(err, common.ERR_INTERNAL):
		this.AccessLogger.PushError(input.GetContext(), err)
	case err != nil:
		this.PushTrace("reset_general_token_id_in_cache", err.Error(), input.GetContext())
	default:
		this.PushTrace("reset_general_token_id_in_cache", "success", input.GetContext())
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

	err := this.RefreshTokenBlackList.SetWithExpire(
		refreshToken.GetTokenID(), struct{}{}, *refreshToken.GetExpireTime(), ctx,
	)

	if err != nil {

		return err
	}

	return nil
}
