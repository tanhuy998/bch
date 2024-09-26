package usecase

import (
	assignmentServicePort "app/adapter/assignment"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	libCommon "app/lib/common"
	actionResultService "app/service/actionResult"
	"encoding/json"
	"net/http"

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
		CreateAssignmentGroupService assignmentServicePort.ICreateAssignmentGroup
		ActionResult                 actionResultService.IActionResult
	}
)

func (this *CreateAssignmentGroupUseCase) Execute(
	input *requestPresenter.CreateAssignmentGroupRequest,
	output *responsePresenter.CreateAssignmentGroupResponse,
) (mvc.Result, error) {

	data := input.Data

	data.CreatedBy = libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID())

	ret, err := this.CreateAssignmentGroupService.Serve(input.AssignmentUUID, data, input.GetContext())

	if err == assignmentServicePort.ERR_ASSIGNMENT_NOT_FOUND {

		output.Message = err.Error()
		raw_content, _ := json.Marshal(output)

		return this.ActionResult.Prepare().
			SetCode(http.StatusNotFound).
			SetContent(raw_content).
			Done()
	}

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Data = ret
	output.Message = "success"

	return this.ActionResult.ServeResponse(output)
}
