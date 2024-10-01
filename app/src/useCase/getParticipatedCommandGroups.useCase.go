package usecase

import (
	authServiceAdapter "app/adapter/auth"
	"app/adapter/responsePresetPort"
	"app/domain/model"
	libCommon "app/lib/common"
	actionResultService "app/service/actionResult"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetParticipatedCommandGroups interface {
		Execute(
			input *requestPresenter.GetParticipatedGroups,
			output *responsePresenter.GetParticipatedGroups,
		) (mvc.Result, error)
	}

	GetParticipatedCommandGroupsUseCase struct {
		GetParticipatedCommandGroups authServiceAdapter.IGetParticipatedCommandGroups // authService.IGetParticipatedCommandGroups
		ResponsePreset               responsePresetPort.IResponsePreset
		ActionResult                 actionResultService.IActionResult
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
