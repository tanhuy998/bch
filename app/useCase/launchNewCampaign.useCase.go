package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	libCommon "app/lib/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	"time"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ILaunchNewCampaign interface {
		Execute(*requestPresenter.LaunchNewCampaignRequest, responsePresenter.ILaunchNewCampaignResponsePresenter) (mvc.Result, error)
	}

	LaunchNewCampaignUseCase struct {
		LaunchNewCampaignService adminService.ILaunchNewCampaign
		ActionResultService      actionResultService.IActionResult
	}
)

func (this *LaunchNewCampaignUseCase) Execute(
	input *requestPresenter.LaunchNewCampaignRequest,
	output responsePresenter.ILaunchNewCampaignResponsePresenter,
) (mvc.Result, error) {

	input.Data.UUID = nil
	input.Data.IssueTime = libCommon.PointerPrimitive(time.Now())

	var inputCampaign *model.Campaign = input.Data

	if inputCampaign == nil {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	createdUUID, err := this.LaunchNewCampaignService.Execute(inputCampaign)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.SetMessage("success")
	output.GetData().CreatedUUID = createdUUID

	return this.ActionResultService.
		Prepare().
		SetCode(201).
		ServeResponse(output)
}
