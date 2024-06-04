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
	IGetPendingCampaigns interface {
		Execute(
			input *requestPresenter.GetPendingCampaignRequest,
			output *responsePresenter.GetPendingCampaingsResponse,
		) (mvc.Result, error)
	}

	GetPendingCampaignsUseCase struct {
		GetPendingCampaignsService adminService.IGetPendingCampaigns
		ActionResultService        actionResultService.IActionResult
	}
)

func (this *GetPendingCampaignsUseCase) Execute(
	input *requestPresenter.GetPendingCampaignRequest,
	output *responsePresenter.GetPendingCampaingsResponse,
) (mvc.Result, error) {

	if input == nil {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	_, err := this.GetPendingCampaignsService.Serve(input.PivotID, input.PageSizeLimit, input.IsPrev)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.Message = "success"

	return this.ActionResultService.ServeResponse(output)
}
