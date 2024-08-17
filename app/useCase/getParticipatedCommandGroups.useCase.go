package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"

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
		GetParticipatedCommandGroups authService.IGetParticipatedCommandGroups
		ActionResult                 actionResultService.IActionResult
	}
)

func (this *GetParticipatedCommandGroupsUseCase) Execute(
	input *requestPresenter.GetParticipatedGroups,
	output *responsePresenter.GetParticipatedGroups,
) (mvc.Result, error) {

	report, err := this.GetParticipatedCommandGroups.Serve(input.UserUUID)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Data = report

	return this.ActionResult.ServeResponse(output)
}
