package usecase

import (
	"app/internal/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"fmt"

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
		CandidateSigningReportService      adminService.ICandidateSigningReport
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

	signingReport, err := this.CandidateSigningReportService.Serve(input.CampaignUUID)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	dataPack, err := this.GetCampaignSignedCandidatesService.Serve(
		input.CampaignUUID, input.PivotID, input.PageSizeLimit, input.IsPrev,
	)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.Message = "sucess"
	output.Data = dataPack.Data

	dataPack.Count = signingReport.SignedCount
	output.CandidateSignedCount = signingReport.SignedCount

	pageNumber := common.PaginationPage(1)
	err = preparePaginationNavigation(output, dataPack, pageNumber)

	if err != nil {
		fmt.Println(3, err)
		return this.ActionResultService.ServeErrorResponse(err)
	}

	return this.ActionResultService.ServeResponse(output)
}
