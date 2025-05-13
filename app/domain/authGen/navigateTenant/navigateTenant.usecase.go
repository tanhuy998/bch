package navigateTenantDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authGenServicePort "app/port/authGenService"
	generalTokenClientServicePort "app/port/generalTokenClient"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"
	"errors"
	"fmt"
)

var (
	errGeneralTokenMustBeRemovedFromClient = fmt.Errorf("general token must be removed from client")
	errLoginAgain                          = errors.Join(common.ERR_UNAUTHORIZED, fmt.Errorf("login again"))
)

type (
	NavigateTenantUseCase struct {
		unitOfWork.GenericUseCase[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant[model.Tenant]]
		unitOfWork.UseCaseResultWrapper[requestPresenter.AuthNavigateTenant, responsePresenter.AuthNavigateTenant[model.Tenant]]
		unitOfWork.UserSessionManipulationUseCase[requestPresenter.AuthNavigateTenant]
		GeneralTokenClient    generalTokenClientServicePort.IGeneralTokenClient
		NavigateTenantService authGenServicePort.INavigateTenant
	}
)

func (this *NavigateTenantUseCase) Execute(
	input *requestPresenter.AuthNavigateTenant,
) (output *responsePresenter.AuthNavigateTenant[model.Tenant], err error) {

	defer this.WrapResults(input, &output, &err)

	generalToken, err := this.GeneralTokenClient.Read(input.GetContext())

	defer func() {

		switch {
		case errors.Is(err, errGeneralTokenMustBeRemovedFromClient):
			err = this.GeneralTokenClient.Remove(input.Context)
			err = libCommon.Ternary(err == nil, common.ERR_UNAUTHORIZED, err)
			return
		case err != nil:
			return
		}
	}()

	switch {
	case err != nil:
		return nil, err
	case generalToken == nil:
		return nil, errLoginAgain
	case generalToken.Expire():
		return nil, errGeneralTokenMustBeRemovedFromClient
	}

	hasUserSession, err := this.HasUserSession(generalToken, input.GetContext())

	if err != nil {

		return nil, err
	}

	if hasUserSession {

		return nil, errGeneralTokenMustBeRemovedFromClient
	}

	// data, err := this.NavigateTenantService.Serve(generalToken.GetUserUUID(), input.GetContext())

	data, err := this.NavigateTenantService.Serve(
		domain_input{generalToken.GetUserUUID(), input.GetContext()},
	)

	if err != nil {

		return nil, err
	}

	output = this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}
