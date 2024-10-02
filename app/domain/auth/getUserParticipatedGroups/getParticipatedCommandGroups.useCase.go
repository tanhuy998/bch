package getUserParticipatedCommandGroupDomain

import (
	libCommon "app/internal/lib/common"
	"app/model"
	actionResultServicePort "app/port/actionResult"
	authServicePort "app/port/auth"
	"app/port/responsePresetPort"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	// IGetParticipatedCommandGroups interface {
	// 	Execute(
	// 		input *requestPresenter.GetParticipatedGroups,
	// 		output *responsePresenter.GetParticipatedGroups,
	// 	) (mvc.Result, error)
	// }

	GetParticipatedCommandGroupsUseCase struct {
		GetParticipatedCommandGroups authServicePort.IGetParticipatedCommandGroups // authService.IGetParticipatedCommandGroups
		ResponsePreset               responsePresetPort.IResponsePreset
		ActionResult                 actionResultServicePort.IActionResult
	}
)

func (this *GetParticipatedCommandGroupsUseCase) Execute(
	input *requestPresenter.GetParticipatedGroups,
	output *responsePresenter.GetParticipatedGroups,
) (mvc.Result, error) {

	//report, err := this.GetParticipatedCommandGroups.Serve(input.UserUUID)

	report, err := this.GetParticipatedCommandGroups.SearchAndRetrieveByModel(
		&model.User{
			UUID:       input.UserUUID,
			TenantUUID: libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID()),
		},
		input.GetContext(),
	)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	if report == nil || len(report.Details) == 0 {

		return this.ResponsePreset.UnAuthorizedResource()
	}

	output.Data = report

	return this.ActionResult.ServeResponse(output)
}
