package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
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

	if input.CampaignUUID == "" {

		return this.ActionResultService.ServeErrorResponse(common.ERR_BAD_REQUEST)
	}

	dataPack, err := this.GetCampaignSignedCandidatesService.Serve(
		input.CampaignUUID, input.PivotID, input.PageSizeLimit, input.IsPrev,
	)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.Message = "sucess"
	output.Data = dataPack.Data

	pageNumber := common.PaginationPage(1)
	err = preparePaginationNavigation(output, dataPack, pageNumber)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	return this.ActionResultService.ServeResponse(output)
}
