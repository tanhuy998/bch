package getTenantCommandGroupDomain

import (
	"app/internal/common"
	"app/model"
	authServicePort "app/port/auth"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"
	"errors"
	"fmt"
)

type (
	GetTenantCommandGroupResponse = responsePresenter.GetTenantCommandGroups[model.CommandGroup]

	GetTenantCommandGroupsUseCase struct {
		unitOfWork.GenericUseCase[requestPresenter.GetTenantCommandGroups, GetTenantCommandGroupResponse]
		GetTenantCommandGroupService authServicePort.IGetTenantCommandGroups[model.CommandGroup]
	}
)

func (this *GetTenantCommandGroupsUseCase) Execute(
	input *requestPresenter.GetTenantCommandGroups,
) (output *GetTenantCommandGroupResponse, err error) {

	defer this.WrapResults(input, &output, &err)

	if !input.IsValidTenantUUID() {

		err = errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("invalid tenant"),
		)
		return
	}

	data, err := this.GetTenantCommandGroupService.Serve(
		input.GetTenantUUID(), input, input.GetContext(),
	)

	if err != nil {

		return
	}

	output = this.GenerateOutput()
	output.Message = "success"
	output.SetData(data)
	//output.ResolvePaginateNavigator(input.GetPageSize())

	return
}
