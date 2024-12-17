package getAssignmentsDomain

import (
	"app/internal/common"
	assignmentServicePort "app/port/assignment"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/unitOfWork"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetAssignmentUseCase struct {
		//usecasePort.UseCase[requestPresenter.GetAssignments, responsePresenter.GetAssignments]
		unitOfWork.GenericUseCase[requestPresenter.GetAssignments, responsePresenter.GetAssignments]
		GetAssignmentService assignmentServicePort.IGetAssignments[primitive.ObjectID]
	}
)

func (this *GetAssignmentUseCase) Execute(
	input *requestPresenter.GetAssignments,
) (output *responsePresenter.GetAssignments, err error) {

	defer func() {

		//log.Default().Println(this.AccessLogger)

		//pthis.PushTrace("err_check", libCommon.Ternary(err == nil, "no error", err.Error()), input.GetContext())

		this.WrapResults(input, &output, &err)
	}()

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	if !input.GetAuthority().IsTenantAgent() {

		return nil, this.ErrorWithContext(
			input, common.ERR_FORBIDEN,
		)
	}

	data, err := this.GetAssignmentService.Serve(
		input.GetTenantUUID(), input, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	output = this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}
