package switchTenantDomain

import (
	"app/internal/common"
	"app/internal/generalToken"
	libCommon "app/internal/lib/common"
	accessTokenServicePort "app/port/accessToken"
	authServicePort "app/port/auth"
	authSignaturesServicePort "app/port/authSignatures"
	generalTokenClientServicePort "app/port/generalTokenClient"
	refreshTokenServicePort "app/port/refreshToken"
	refreshTokenClientPort "app/port/refreshTokenClient"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	errGeneralTokenMustBeRemovedFromClient = fmt.Errorf("general token must be removed from client")
)

type (
	SwitchTenantUseCase struct {
		unitOfWork.MongoUserSessionCacheUseCase[responsePresenter.SwitchTenant]
		unitOfWork.GenericUseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
		unitOfWork.UseCaseResultWrapper[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
		unitOfWork.OperationLogger
		GeneralTokenClientService generalTokenClientServicePort.IGeneralTokenClient
		SwitchTenantService       authSignaturesServicePort.ISwitchTenant
		AccessTokenManipulator    accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenClientService refreshTokenClientPort.IRefreshTokenClient
		RefreshTokenIDProvider    refreshTokenIdServicePort.IRefreshTokenIDProvider
		GetSingleUserService      authServicePort.IGetSingleUser
	}
)

func (this *SwitchTenantUseCase) Execute(
	input *requestPresenter.SwitchTenant,
) (output *responsePresenter.SwitchTenant, err error) {

	defer this.WrapResults(input, &output, &err)

	generalToken, err := this.GeneralTokenClientService.Read(input.GetContext())

	if err != nil {

		return nil, err
	}

	if generalToken == nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	this.DebugLogger.Push(
		"show_general_token_id", generalToken.GetTokenID().String(), input.GetContext(),
	)

	defer func() {

		switch {
		case errors.Is(err, errGeneralTokenMustBeRemovedFromClient):
			err = this.GeneralTokenClientService.Remove(input.GetContext())
			err = libCommon.Ternary(err == nil, common.ERR_UNAUTHORIZED, err)
			return
		case err != nil:
			return
		}

		if output == nil {

			return
		}

		output.Data.User, _ = this.GetSingleUserService.Serve(
			generalToken.GetUserUUID(), input.GetContext(),
		)
	}()

	switch generalTokenIdInCache, err := this.GeneralTokenWhiteList.Has(generalToken.GetTokenID(), input.GetContext()); {
	case err != nil:
		return nil, err
	case generalToken.Expire() || generalTokenIdInCache:
		return nil, errGeneralTokenMustBeRemovedFromClient
	}

	output, err = this.ModifyUserSession(
		input.GetContext(),
		func(sesstionCtx context.Context) (ret *responsePresenter.SwitchTenant, err error) {

			if _, ok := sesstionCtx.(mongo.SessionContext); ok {

				this.DebugLogger.Push("check_session_context", "true", input.GetContext())
			}

			at, rt, err := this.SwitchTenantService.Serve(*input.TenantUUID, generalToken, sesstionCtx)

			if errors.Is(err, common.ERR_UNAUTHORIZED) {

				e := this.RefreshTokenClientService.Remove(input.GetContext())

				if e != nil {

					return nil, e
				}

				return nil, err
			}

			if err != nil {

				return nil, err
			}

			defer func() {

				if err != nil {

					ret = nil
					return
				}

				err = this.GeneralTokenClientService.Remove(input.GetContext())

				if err != nil {

					ret = nil
					return //nil, err
				}

				err = this.RefreshTokenClientService.Write(input.GetContext(), rt)

				if err != nil {

					ret = nil
					return //nil, err
				}
			}()

			at_str, err := this.AccessTokenManipulator.SignString(at)

			if err != nil {

				return nil, err
			}

			err = this.manageSessions(generalToken.GetTokenID(), rt, sesstionCtx)

			if err != nil {

				return nil, err
			}

			output := this.GenerateOutput()
			output.Data.AccessToken = at_str

			return output, nil
		},
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	return output, nil
}

func (this *SwitchTenantUseCase) manageSessions(
	generalTokenID generalToken.GeneralTokenID, newRefreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context,
) (err error) {

	err = this.RemoveUserSession(
		ctx,
		//generalToken.GetUserUUID(),
		newRefreshToken.GetUserUUID(),
	)

	if err != nil {

		return
	}

	expire := newRefreshToken.GetExpireTime()

	if expire != nil {

		err = this.GeneralTokenWhiteList.SetWithExpire(
			generalTokenID, struct{}{}, *expire, ctx,
		)
	} else {

		_, err = this.GeneralTokenWhiteList.Set(
			generalTokenID, struct{}{}, ctx,
		)
	}

	if err != nil {

		return err
	}

	return nil

}
