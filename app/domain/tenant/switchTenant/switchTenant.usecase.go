package switchTenantDomain

import (
	"app/internal/common"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	generalTokenServicePort "app/port/generalToken"
	generalTokenClientServicePort "app/port/generalTokenClient"
	refreshTokenClientPort "app/port/refreshTokenClient"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	tenantServicePort "app/port/tenant"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	SwitchTenantUseCase struct {
		usecasePort.UserSessionCacheUseCase
		GeneralTokenClientService generalTokenClientServicePort.IGeneralTokenClient
		SwitchTenantService       tenantServicePort.ISwitchTenant
		AccessTokenManipulator    accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenClientService refreshTokenClientPort.IRefreshTokenClient
		RefreshTokenIDProvider    refreshTokenIdServicePort.IRefreshTokenIDProvider
		usecasePort.UseCase[requestPresenter.SwitchTenant, responsePresenter.SwitchTenant]
	}
)

func (this *SwitchTenantUseCase) Execute(
	input *requestPresenter.SwitchTenant,
) (*responsePresenter.SwitchTenant, error) {

	generalToken, err := this.GeneralTokenClientService.Read(input.GetContext())

	if err != nil {

		return nil, err
	}

	if generalToken == nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	session, err := this.MongoClient.StartSession()

	if err != nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(err),
		)
	}

	defer session.EndSession(input.GetContext())

	unknown, err := session.WithTransaction(
		input.GetContext(),
		func(sesstionCtx mongo.SessionContext) (ret interface{}, err error) {

			at, rt, err := this.SwitchTenantService.Serve(*input.TenantUUID, generalToken, sesstionCtx)
			fmt.Println(1)

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
				fmt.Println(2)
				if err != nil {

					ret = nil
					return //nil, err
				}
			}()

			at_str, err := this.AccessTokenManipulator.SignString(at)
			fmt.Println(3)
			if err != nil {

				return nil, err
			}
			fmt.Println(4)

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

	output, ok := unknown.(*responsePresenter.SwitchTenant)

	if !ok {

		return nil, libError.NewInternal(fmt.Errorf("unknown error"))
	}

	return output, nil
}

func (this *SwitchTenantUseCase) manageSessions(
	generalToken generalTokenServicePort.IGeneralToken, ctx context.Context,
) (err error) {

	userSessions, err := this.UserSessionRepo.FindMany(
		bson.D{
			{"userUUID", generalToken.GetUserUUID()},
		},
		ctx,
	)

	if err != nil {

		return err
	}

	defer func() {

		if err != nil {

			return
		}

		for _, v := range userSessions {
			// Delete caches of current user sessions
			// ctx of this funciton is a transaction context, therefore fetched data from
			// db have not committed until the whole transaction committed
			_, err = this.GeneralTokenWhiteList.Delete(*v.SessionID, ctx)

			if err != nil {

				return
			}
		}
	}()

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
