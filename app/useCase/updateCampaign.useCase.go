package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IUpdateCampaign interface {
		Execute(*requestPresenter.UpdateCampaignRequest, *responsePresenter.UpdateCampaignResponse) (mvc.Result, error)
	}
)
