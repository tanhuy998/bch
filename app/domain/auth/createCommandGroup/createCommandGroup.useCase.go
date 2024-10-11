package createCommandGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"fmt"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateCommandGroup interface {
		Execute(
			input *requestPresenter.CreateCommandGroupRequest,
			output *responsePresenter.CreateCommandGroupResponse,
		) (mvc.Result, error)
	}

	CreateCommandGroupUseCase struct {
		usecasePort.UseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse]
		CreateCommandGroupService authServicePort.ICreateCommandGroup
	}
)

func (this *CreateCommandGroupUseCase) Execute(
	input *requestPresenter.CreateCommandGroupRequest,
) (*responsePresenter.CreateCommandGroupResponse, error) {

	data := input.Data

	authority := input.GetAuthority()

	if authority == nil ||
		!input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	data.CreatedBy = libCommon.PointerPrimitive(authority.GetUserUUID())

	data, err := this.CreateCommandGroupService.CreateByModel(input.GetTenantUUID(), data, input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	if data == nil {

		return nil, this.ErrorWithContext(
			input, libError.NewInternal(fmt.Errorf("cannot create command group , try again")),
		)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data.UUID = *data.UUID

	return output, nil
}
