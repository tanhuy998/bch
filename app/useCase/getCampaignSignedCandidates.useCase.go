package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetCampaignSignedCandidates interface {
		Execute(
			input *requestPresenter.GetCampaignSignedCandidatesRequest,
			output *responsePresenter.GetCampaignSignedCandidatesResponse,
		) (mvc.Result, error)
	}

	GetCampaignSignedCandidatesUseCase struct {
		GetCampaignSignedCandidatesService adminService.IGetCampaignSignedCandidates
		ActionResultService                actionResultService.IActionResult
	}
)

func (this GetCampaignSignedCandidatesUseCase) Execute(
	input *requestPresenter.GetCampaignSignedCandidatesRequest,
	output *responsePresenter.GetCampaignSignedCandidatesResponse,
) (mvc.Result, error) {

	res, err := this.GetCampaignSignedCandidatesService.Serve(
		input.CampaignUUID, input.PivotID, input.PageSizeLimit, input.IsPrev,
	)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

}
