package reportUserParticipatedCommandGroupsDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	ReportParticipatedCommandGroupsUseCase struct {
		usecasePort.UseCase[requestPresenter.ReportParticipatedGroups, responsePresenter.ReportParticipatedGroups]
		ReportParticipatedCommandGroups authServicePort.IReportParticipatedCommandGroups // authService.IGetParticipatedCommandGroups
	}
)

func (this *ReportParticipatedCommandGroupsUseCase) Execute(
	input *requestPresenter.ReportParticipatedGroups,
) (*responsePresenter.ReportParticipatedGroups, error) {

	report, err := this.ReportParticipatedCommandGroups.Serve(
		input.GetTenantUUID(), *input.UserUUID, input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	if report == nil {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = report

	return output, nil
}
