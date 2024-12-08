package getTenantCommandGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	paginateServicePort "app/port/paginate"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetTenantCommandGroupResponse = responsePresenter.GetTenantCommandGroups[model.CommandGroup]

	GetTenantCommandGroupsUseCase struct {
		unitOfWork.GenericUseCase[requestPresenter.GetTenantCommandGroups, GetTenantCommandGroupResponse]
		GetTenantCommandGroupService paginateServicePort.IPaginate[model.CommandGroup, primitive.ObjectID] // authServicePort.IGetTenantCommandGroup[model.CommandGroup]
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

	data, err := this.GetTenantCommandGroupService.Paginate(
		input.GetTenantUUID(), input.PageNumber, input.PageSize, libCommon.PointerPrimitive(input.GetCursor()), input.IsPrev, input.GetContext(),
	)

	if err != nil {

		return
	}

	output = this.GenerateOutput()
	output.Message = "success"
	output.SetData(data)
	output.ResolvePaginateNavigator(input.GetPageSize())

	return
}
