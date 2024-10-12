package createAssignmentGroupDomain

import (
	libCommon "app/internal/lib/common"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateAssignmentGroup interface {
		Execute(
			input *requestPresenter.CreateAssignmentGroupRequest,
			output *responsePresenter.CreateAssignmentGroupResponse,
		) (mvc.Result, error)
	}

	CreateAssignmentGroupUseCase struct {
		usecasePort.UseCase[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse]
		CreateAssignmentGroupService assignmentServicePort.ICreateAssignmentGroup
		//ActionResult                 actionResultServicePort.IActionResult
	}
)

func (this *CreateAssignmentGroupUseCase) Execute(
	input *requestPresenter.CreateAssignmentGroupRequest,
) (*responsePresenter.CreateAssignmentGroupResponse, error) {

	data := input.Data

	data.CreatedBy = libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID())

	ret, err := this.CreateAssignmentGroupService.Serve(
		input.GetTenantUUID(), *input.AssignmentUUID, *input.CommandGroupUUID, data, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Data = ret
	output.Message = "success"

	return output, nil
}
