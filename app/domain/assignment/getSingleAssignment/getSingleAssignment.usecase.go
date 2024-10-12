package getSingleAssignmentDomain

import (
	"app/internal/common"
	assignmentServicePort "app/port/assignment"
	"app/port/responsePresetPort"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

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
		usecasePort.UseCase[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse]
		GetSingleAssignmnetService assignmentServicePort.IGetSingleAssignnment
		ResponsePreset             responsePresetPort.IResponsePreset
	}
)

func (this *GetSingleAssignmentUseCase) Execute(
	input *requestPresenter.GetSingleAssignmentRequest,
) (*responsePresenter.GetSingleAssignmentResponse, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	data, err := this.GetSingleAssignmnetService.Serve(
		input.GetTenantUUID(), *input.UUID, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Data = data
	output.Message = "success"

	return output, nil
}
