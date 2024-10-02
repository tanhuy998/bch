package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleCampaign interface {
		Execute(
			*requestPresenter.GetSingleCampaignRequest,
			*responsePresenter.GetSingleCampaignResponse,
		) (mvc.Result, error)
	}
)
