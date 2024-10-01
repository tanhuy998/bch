package createCommandGroupDomain

import (
	libCommon "app/src/internal/lib/common"
	actionResultServicePort "app/src/port/actionResult"
	authServicePort "app/src/port/auth"
	"app/src/port/responsePresetPort"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"encoding/json"

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
		CreateCommandGroupService authServicePort.ICreateCommandGroup //authService.ICreateCommandGroup
		ResponsePreset            responsePresetPort.IResponsePreset
		//GetSingleCommandGroupService authService.IGetSingleCommandGroup
		ActionResult actionResultServicePort.IActionResult
	}
)

func (this *CreateCommandGroupUseCase) Execute(
	input *requestPresenter.CreateCommandGroupRequest,
	output *responsePresenter.CreateCommandGroupResponse,
) (mvc.Result, error) {

	data := input.Data
	data.CreatedBy = libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID())
	data.TenantUUID = libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID())

	data, err := this.CreateCommandGroupService.CreateByModel(data, input.GetContext())

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	if data == nil {

		return this.ResponsePreset.InternalError()
	}

	output.Message = "success"
	output.Data.UUID = *data.UUID

	rawContent, _ := json.Marshal(output)

	return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent).Done()
}
