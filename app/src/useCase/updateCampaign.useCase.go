package usecase

import (
	"app/domain/model"
	"app/internal/common"
	libCommon "app/lib/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IUpdateCampaign interface {
		Execute(*requestPresenter.UpdateCampaignRequest, *responsePresenter.UpdateCampaignResponse) (mvc.Result, error)
	}

	UpdateCampaignUseCase struct {
		GetSingleCampaignService adminService.IGetCampaign
		ModifyCampaignService    adminService.IModifyExistingCampaign
		ActionResultService      actionResultService.IActionResult
	}
)

func (this *UpdateCampaignUseCase) Execute(
	input *requestPresenter.UpdateCampaignRequest,
	output *responsePresenter.UpdateCampaignResponse,
) (mvc.Result, error) {

	var (
		uuid            string          = input.UUID
		campaignUpdated *model.Campaign = new(model.Campaign)
	)

	inputModel := input.Data

	if libCommon.Or(uuid == "", inputModel == nil) {

		return this.ActionResultService.Prepare().ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	targetCampaingn, err := this.GetSingleCampaignService.Serve(uuid)

	if err != nil || targetCampaingn == nil {

		return this.ActionResultService.Prepare().ServeErrorResponse(common.ERR_HTTP_NOT_FOUND)
	}

	if inputModel.Title != nil || *inputModel.Title != "" {

		campaignUpdated.Title = inputModel.Title

	} else {

		campaignUpdated.Title = nil
	}

	campaignUpdated.Expire = inputModel.Expire

	err = this.ModifyCampaignService.Serve(uuid, campaignUpdated)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.Message = "success"

	return this.ActionResultService.Prepare().ServeResponse(output)
}
