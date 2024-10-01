package usecase

import (
	"app/internal/common"
	libCommon "app/lib/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleCampaign interface {
		Execute(
			*requestPresenter.GetSingleCampaignRequest,
			*responsePresenter.GetSingleCampaignResponse,
		) (mvc.Result, error)
	}

	GetSingleCampaignUseCase struct {
		GetSingleCampaignService adminService.IGetCampaign
		ActionResult             actionResultService.IActionResult
	}
)

func (this *GetSingleCampaignUseCase) Execute(
	input *requestPresenter.GetSingleCampaignRequest,
	output *responsePresenter.GetSingleCampaignResponse,
) (mvc.Result, error) {

	if libCommon.Or(input == nil, input.UUID == "") {

		return this.ActionResult.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	data, err := this.GetSingleCampaignService.Serve(input.UUID)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "succes"
	output.Data = data

	return this.ActionResult.ServeResponse(output)
}
