package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	libCommon "app/lib/common"
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

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	createdUUID, err := this.LaunchNewCampaignService.Execute(inputCampaign)

	if err != nil {

		return nil, err
	}

	res := NewResponse()

	output.SetMessage("success")
	output.GetData().CreatedUUID = createdUUID
	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	res.Code = 201

	return res, nil
}
