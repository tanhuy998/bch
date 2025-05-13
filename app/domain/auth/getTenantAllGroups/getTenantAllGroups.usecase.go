package getTenantAllGroupsDomain

import (
	"app/internal/common"
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"
	"fmt"
)

type (
	GetTenantAllGroupUseCase struct {
		usecasePort.UseCase[requestPresenter.GetTenantAllGroups, responsePresenter.GetTenantAllGroups[model.CommandGroup]]
		GetTenantAllGroupsService authServicePort.IGetTenantAllGroups
	}
)

func (this *GetTenantAllGroupUseCase) Execute(
	input *requestPresenter.GetTenantAllGroups,
) (*responsePresenter.GetTenantAllGroups[model.CommandGroup], error) {

	auth := input.GetAuthority()

	switch {
	case !input.IsValidTenantUUID():
		return nil, common.ERR_BAD_REQUEST
	case auth.GetTenantUUID() != input.GetTenantUUID():
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("the current user is not in tenant"))
	}

	//data, err := this.GetTenantAllGroupsService.Serve(input.GetTenantUUID(), input.GetContext())
	data, err := this.GetTenantAllGroupsService.Serve(input)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = data

	return output, nil
}
