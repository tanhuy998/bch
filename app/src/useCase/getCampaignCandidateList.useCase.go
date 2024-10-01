package usecase

import (
	"app/domain/model"
	"app/internal/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetCampaignCandidateList interface {
		Execute(
			*requestPresenter.GetCampaignCandidateListRequest,
			*responsePresenter.GetCampaignCandidateListResponse,
		) (mvc.Result, error)
	}

	GetCampaignCandidateListUseCase struct {
		GetCampaignCandidateListService adminService.IGetCampaignCandidateList
		ActionResultService             actionResultService.IActionResult
	}
)

func (this *GetCampaignCandidateListUseCase) Execute(
	input *requestPresenter.GetCampaignCandidateListRequest,
	output *responsePresenter.GetCampaignCandidateListResponse,
) (mvc.Result, error) {

	var (
		candidatePivot_id string = input.PivotID
		err               error
	)

	dataPack, err := this.GetCampaignCandidateListService.Serve(
		input.CampaignUUID,
		candidatePivot_id,
		input.PageSizeLimit,
		input.IsPrev,
	)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.Message = "success"
	output.Data = dataPack.Data

	if input.ExposeHeader {

		output.Header.Campaign = dataPack.Header
	}

	pageNumber := common.PAGINATION_FIRST_PAGE

	err = preparePaginationNavigation[model.Candidate](output, dataPack.PaginationPack, pageNumber)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	return this.ActionResultService.ServeResponse(output)
}
