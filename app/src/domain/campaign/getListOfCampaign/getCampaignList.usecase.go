package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	"app/repository"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	RetrievedData_T = repository.PaginationPack[model.Campaign]

	IGetCampaignList interface {
		Execute(
			input *requestPresenter.GetCampaignListRequest,
			output *responsePresenter.GetCampaignListResponse,
		) (mvc.Result, error)
	}

	GetCampaignListUseCase struct {
		GetCampaignListService adminService.IGetCampaignList
		ActionResultService    actionResultService.IActionResult
	}
)

func (this *GetCampaignListUseCase) Execute(
	input *requestPresenter.GetCampaignListRequest,
	output *responsePresenter.GetCampaignListResponse,
) (mvc.Result, error) {

	if input == nil {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	dataPack, err := this.GetCampaignListService.Serve(input.PivotID, input.PageSizeLimit, input.IsPrev)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	pageNumber := common.PaginationPage(1)

	output.Message = "success"
	output.Data = dataPack.Data

	err = preparePaginationNavigation[model.Campaign](output, dataPack, pageNumber)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	return this.ActionResultService.ServeResponse(output)
}
