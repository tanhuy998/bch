package getTenantAllGroupsDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GetTenantAllGroupUseCase struct {
		usecasePort.UseCase[requestPresenter.GetTenantAllGroups, responsePresenter.GetTenantAllGroups]
		GetTenantAllGroupsService authServicePort.IGetTenantAllGroups
	}
)

func (this *GetTenantAllGroupUseCase) Execute(
	input *requestPresenter.GetTenantAllGroups,
) (*responsePresenter.GetTenantAllGroups, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	data, err := this.GetTenantAllGroupsService.Serve(input.GetTenantUUID(), input.GetContext())

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = data

	return output, nil
}
