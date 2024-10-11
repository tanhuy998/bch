package getUserParticipatedCommandGroupDomain

import (
	"app/internal/common"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	// IGetParticipatedCommandGroups interface {
	// 	Execute(
	// 		input *requestPresenter.GetParticipatedGroups,
	// 		output *responsePresenter.GetParticipatedGroups,
	// 	) (mvc.Result, error)
	// }

	GetParticipatedCommandGroupsUseCase struct {
		usecasePort.UseCase[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups]
		GetParticipatedCommandGroups authServicePort.IGetParticipatedCommandGroups // authService.IGetParticipatedCommandGroups
	}
)

func (this *GetParticipatedCommandGroupsUseCase) Execute(
	input *requestPresenter.GetParticipatedGroups,
) (*responsePresenter.GetParticipatedGroups, error) {

	//report, err := this.GetParticipatedCommandGroups.Serve(input.UserUUID)

	// report, err := this.GetParticipatedCommandGroups.SearchAndRetrieveByModel(
	// 	&model.User{
	// 		UUID:       input.UserUUID,
	// 		TenantUUID: libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID()),
	// 	},
	// 	input.GetContext(),
	// )

	report, err := this.GetParticipatedCommandGroups.Serve(
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
