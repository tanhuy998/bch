package createAssignmentDomain

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	actionResultServicePort "app/src/port/actionResult"
	assignmentServicePort "app/src/port/assignment"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"time"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateAssignment interface {
		Execute(
			input *requestPresenter.CreateAssigmentRequest,
			output *responsePresenter.CreateAssignmentResponse,
		) (mvc.Result, error)
	}

	CreateAssignmentUseCase struct {
		CreateAssignmentService assignmentServicePort.ICreateAssignment
		ActionResult            actionResultServicePort.IActionResult
	}
)

func (this *CreateAssignmentUseCase) Execute(
	input *requestPresenter.CreateAssigmentRequest,
	output *responsePresenter.CreateAssignmentResponse,
) (mvc.Result, error) {

	inputData := input.Data

	model := &model.Assignment{
		Title:      inputData.Title,
		TenantUUID: libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID()),
		CreatedAt:  libCommon.PointerPrimitive(time.Now()),
		CreatedBy:  libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID()),
	}

	data, err := this.CreateAssignmentService.Serve(model, input.GetContext())

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Data = data

	return this.ActionResult.ServeResponse(output)
}
