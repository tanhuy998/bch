package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetPendingCampaigns interface {
		Execute(
			input *requestPresenter.GetPendingCampaignRequest,
			output *responsePresenter.GetPendingCampaingsResponse,
		) (mvc.Result, error)
	}
)
