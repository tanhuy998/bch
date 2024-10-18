package navigateTenantDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	generalTokenClientServicePort "app/port/generalTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"
	"fmt"
)

type (
	NavigateTenantUseCase struct {
		usecasePort.UseCase[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant]
		GeneralTokenClient    generalTokenClientServicePort.IGeneralTokenClient
		NavigateTenantService authServicePort.INavigateTenant
	}
)

func (this *NavigateTenantUseCase) Execute(
	input *requestPresenter.AuthNavigateTenant,
) (*responsePresenter.AuthNavigateTenant, error) {

	generalToken, err := this.GeneralTokenClient.Read(input.GetContext())

	switch {
	case err != nil:
		return nil, err
	case generalToken == nil:
		fmt.Println(1)
		return nil, errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("login again"))
	case generalToken.Expire():
		fmt.Println(2)
		return nil, errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("login again"))
	}

	data, err := this.NavigateTenantService.Serve(generalToken.GetUserUUID(), input.GetContext())

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}
