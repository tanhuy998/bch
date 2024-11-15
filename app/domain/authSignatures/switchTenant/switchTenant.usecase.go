package switchTenantDomain

import (
	"app/internal/common"
	accessTokenServicePort "app/port/accessToken"
	authServicePort "app/port/auth"
	authSignaturesServicePort "app/port/authSignatures"
	generalTokenServicePort "app/port/generalToken"
	generalTokenClientServicePort "app/port/generalTokenClient"
	refreshTokenClientPort "app/port/refreshTokenClient"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"
	"context"
	"errors"
)

type (
	SwitchTenantUseCase struct {
		unitOfWork.MongoUserSessionCacheUseCase[responsePresenter.SwitchTenant]
		unitOfWork.GenericUseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
		unitOfWork.UseCaseResultWrapper[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
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

	defer func() {

		if err != nil {

			return
		}

		if output == nil {

			return
		}

		output.Data.User, _ = this.GetSingleUserService.Serve(
			generalToken.GetUserUUID(), input.GetContext(),
		)
	}()

	output, err = this.ModifyUserSession(
		input.GetContext(),
		func(sesstionCtx context.Context) (ret *responsePresenter.SwitchTenant, err error) {

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

			err = this.manageSessions(generalToken, sesstionCtx)

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
	generalToken generalTokenServicePort.IGeneralToken, ctx context.Context,
) (err error) {

	err = this.RemoveUserSession(
		ctx,
		generalToken.GetUserUUID(),
	)

	if err != nil {

		return
	}

	expire := generalToken.GetExpiretime()

	if expire != nil {

		err = this.GeneralTokenWhiteList.SetWithExpire(
			generalToken.GetTokenID(), struct{}{}, *expire, ctx,
		)
	} else {

		_, err = this.GeneralTokenWhiteList.Set(
			generalToken.GetTokenID(), struct{}{}, ctx,
		)
	}

	if err != nil {

		return err
	}

	return nil

}
