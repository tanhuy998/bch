package navigateTenantDomain

import (
	"app/internal/common"
	authGenServicePort "app/port/authGenService"
	generalTokenClientServicePort "app/port/generalTokenClient"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"
	"errors"
	"fmt"
)

type (
	NavigateTenantUseCase struct {
		unitOfWork.GenericUseCase[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant]
		unitOfWork.UseCaseResultWrapper[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant]
		GeneralTokenClient    generalTokenClientServicePort.IGeneralTokenClient
		NavigateTenantService authGenServicePort.INavigateTenant
	}
)

func (this *NavigateTenantUseCase) Execute(
	input *requestPresenter.AuthNavigateTenant,
) (output *responsePresenter.AuthNavigateTenant, err error) {

	defer this.WrapResults(input, &output, &err)

	generalToken, err := this.GeneralTokenClient.Read(input.GetContext())

	switch {
	case err != nil:
		return nil, err
	case generalToken == nil:
		return nil, errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("login again"))
	case generalToken.Expire():
		return nil, errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("login again"))
	}

	data, err := this.NavigateTenantService.Serve(generalToken.GetUserUUID(), input.GetContext())

	if err != nil {

		return nil, err
	}

	output = this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}
