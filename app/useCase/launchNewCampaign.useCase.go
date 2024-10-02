package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ILaunchNewCampaign interface {
		Execute(*requestPresenter.LaunchNewCampaignRequest, *responsePresenter.LaunchNewCampaignResponse) (mvc.Result, error)
	}
)
