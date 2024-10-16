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
		usecasePort.AuthDomainUseCase[*requestPresenter.GetSingleAssignmentRequest]
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

	err := this.CheckUserJoinedAssignment(
		input, *input.AssignmentUUID,
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	data, err := this.GetSingleAssignmnetService.Serve(
		input.GetTenantUUID(), *input.AssignmentUUID, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Data = data
	output.Message = "success"

	return output, nil
}

// func (this *GetSingleAssignmentUseCase) validateAuthority(
// 	input *requestPresenter.GetSingleAssignmentRequest,
// ) error {

// 	return nil
// }
