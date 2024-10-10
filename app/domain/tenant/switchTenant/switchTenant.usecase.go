package switchTenantDomain

import (
	"app/internal/common"
	accessTokenServicePort "app/port/accessToken"
	generalTokenClientServicePort "app/port/generalTokenClient"
	refreshTokenClientPort "app/port/refreshTokenClient"
	tenantServicePort "app/port/tenant"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"
	"fmt"
)

type (
	SwitchTenantUseCase struct {
		GeneralTokenClientService generalTokenClientServicePort.IGeneralTokenClient
		SwitchTenantService       tenantServicePort.ISwitchTenant
		AccessTokenManipulator    accessTokenServicePort.IAccessTokenManipulator
		RefreshTokenClientService refreshTokenClientPort.IRefreshTokenClient
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

	at, rt, err := this.SwitchTenantService.Serve(*input.TenantUUID, generalToken, input.GetContext())
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

	err = this.RefreshTokenClientService.Write(input.GetContext(), rt)
	fmt.Println(2)
	if err != nil {

		return nil, err
	}

	at_str, err := this.AccessTokenManipulator.SignString(at)
	fmt.Println(3)
	if err != nil {

		return nil, err
	}
	fmt.Println(4)
	output := this.GenerateOutput()
	output.Data.AccessToken = at_str

	return output, nil
}
