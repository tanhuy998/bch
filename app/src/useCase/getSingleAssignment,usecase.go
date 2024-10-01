package usecase

import (
	assignmentServicePort "app/adapter/assignment"
	"app/adapter/responsePresetPort"
	actionResultService "app/service/actionResult"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleAssignment interface {
		Execute(
			input *requestPresenter.GetSingleAssignmentRequest,
			output *responsePresenter.GetSingleAssignmentResponse,
		) (mvc.Result, error)
	}

	GetSingleAssignmentUseCase struct {
		GetSingleAssignmnetService assignmentServicePort.IGetSingleAssignnment
		ActionResult               actionResultService.IActionResult
		ResponsePreset             responsePresetPort.IResponsePreset
	}
)

func (this *GetSingleAssignmentUseCase) Execute(
	input *requestPresenter.GetSingleAssignmentRequest,
	output *responsePresenter.GetSingleAssignmentResponse,
) (mvc.Result, error) {

	data, err := this.GetSingleAssignmnetService.Serve(input.UUID, input.GetContext())

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	if data == nil {

		return this.ResponsePreset.NotFound()
	}

	if *data.TenantUUID != input.GetAuthority().GetTenantUUID() {

		return this.ResponsePreset.UnAuthorizedResource()
	}

	output.Data = data
	output.Message = "success"

	return this.ActionResult.ServeResponse(output)
}
