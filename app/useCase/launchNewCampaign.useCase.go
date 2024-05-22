package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ILaunchNewCampaign interface {
		Execute(*requestPresenter.LaunchNewCampaignRequest, *responsePresenter.LaunchNewCampaignResponse) (mvc.Result, error)
	}

	LaunchNewCampaignUseCase struct {
		LaunchNewCampaignService adminService.ILaunchNewCampaign
	}
)

func (this *LaunchNewCampaignUseCase) Execute(
	input *requestPresenter.LaunchNewCampaignRequest,
	output *responsePresenter.LaunchNewCampaignResponse,
) (mvc.Result, error) {

	var inputCampaign *model.Campaign = input.Data

	if inputCampaign == nil {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	err := this.LaunchNewCampaignService.Execute(inputCampaign)

	if err != nil {

		return nil, err
	}

	res := newResponse()

	output.Message = "success"
	err = marshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	res.Code = 201

	return res, nil
}
