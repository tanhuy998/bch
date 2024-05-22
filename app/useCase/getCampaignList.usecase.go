package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetCampaignList interface {
		Execute(
			input *requestPresenter.GetCampaignListRequest,
			output *responsePresenter.GetCampaignListResponse,
		) (mvc.Result, error)
	}

	GetCampaignListUseCase struct {
		GetCampaignListService adminService.IGetCampaignList
	}
)

func (this *GetCampaignListUseCase) Execute(
	input *requestPresenter.GetCampaignListRequest,
	output *responsePresenter.GetCampaignListResponse,
) (mvc.Result, error) {

	if input == nil {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	_, err := this.GetCampaignListService.Execute(input.PageNumber)

	if err != nil {

		return nil, err
	}

	res := newResponse()
	output.Message = "success"

	err = marshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}
