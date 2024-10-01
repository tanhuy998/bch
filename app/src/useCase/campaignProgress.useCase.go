package usecase

import (
	"app/internal/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICampaignProgress interface {
		Execute(
			input *requestPresenter.CampaignProgressRequestPresenter,
			output *responsePresenter.CampaignProgressResponsePresenter,
		) (mvc.Result, error)
	}

	CampaignProgressUseCase struct {
		CampaignProgressService adminService.ICandidateSigningReport
		ActionResultService     actionResultService.IActionResult
	}
)

func (this *CampaignProgressUseCase) Execute(
	input *requestPresenter.CampaignProgressRequestPresenter,
	output *responsePresenter.CampaignProgressResponsePresenter,
) (mvc.Result, error) {

	if input.CampaignUUID == "" {

		return this.ActionResultService.ServeErrorResponse(common.ERR_BAD_REQUEST)
	}

	signingReport, err := this.CampaignProgressService.Serve(input.CampaignUUID)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.Data = *signingReport

	return this.ActionResultService.ServeResponse(output)
}
