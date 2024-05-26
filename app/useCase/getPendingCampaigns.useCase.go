package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
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
	}
)

func (this *GetPendingCampaignsUseCase) Execute(
	input *requestPresenter.GetPendingCampaignRequest,
	output *responsePresenter.GetPendingCampaingsResponse,
) (mvc.Result, error) {

	if input == nil {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	_, err := this.GetPendingCampaignsService.Serve(input.PivotID, input.PageSizeLimit, input.IsPrev)

	if err != nil {

		return nil, err
	}

	res := NewResponse()
	output.Message = "success"

	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}
