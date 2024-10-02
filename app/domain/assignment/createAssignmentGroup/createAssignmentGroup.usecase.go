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
		usecasePort.UseCase[responsePresenter.CreateAssignmentGroupResponse]
		CreateAssignmentGroupService assignmentServicePort.ICreateAssignmentGroup
		//ActionResult                 actionResultServicePort.IActionResult
	}
)

func (this *CreateAssignmentGroupUseCase) Execute(
	input *requestPresenter.CreateAssignmentGroupRequest,
) (*responsePresenter.CreateAssignmentGroupResponse, error) {

	data := input.Data

	data.CreatedBy = libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID())

	ret, err := this.CreateAssignmentGroupService.Serve(input.AssignmentUUID, data, input.GetContext())

	// if err == assignmentServicePort.ERR_ASSIGNMENT_NOT_FOUND {

	// 	output.Message = err.Error()
	// 	raw_content, _ := json.Marshal(output)

	// 	return this.ActionResult.Prepare().
	// 		SetCode(http.StatusNotFound).
	// 		SetContent(raw_content).
	// 		Done()
	// }

	if err != nil {

		//return this.ActionResult.ServeErrorResponse(err)

		return nil, err
	}

	output := this.GenerateOutput()

	output.Data = ret
	output.Message = "success"

	//return this.ActionResult.ServeResponse(output)

	return output, nil
}
