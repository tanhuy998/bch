package createAssignmentDomain

import (
	libCommon "app/internal/lib/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
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
		usecasePort.UseCase[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse]
		CreateAssignmentService assignmentServicePort.ICreateAssignment
		//ActionResult            actionResultServicePort.IActionResult
	}
)

func (this *CreateAssignmentUseCase) Execute(
	input *requestPresenter.CreateAssigmentRequest,
) (*responsePresenter.CreateAssignmentResponse, error) {

	inputData := input.Data

	model := &model.Assignment{
		Title:      inputData.Title,
		TenantUUID: libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID()),
		CreatedAt:  libCommon.PointerPrimitive(time.Now()),
		CreatedBy:  libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID()),
	}

	data, err := this.CreateAssignmentService.Serve(model, input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
		//return this.ActionResult.ServeErrorResponse(err)
	}

	// output.Data = data

	// return this.ActionResult.ServeResponse(output)

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = data

	return output, nil
}
