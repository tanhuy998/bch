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

	if err != nil {

		return nil, err
	}

	err = this.RefreshTokenClientService.Write(input.GetContext(), rt)

	if err != nil {

		return nil, err
	}

	at_str, err := this.AccessTokenManipulator.SignString(at)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()
	output.Data.AccessToken = at_str

	return output, nil
}
