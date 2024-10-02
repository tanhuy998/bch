package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICampaignProgress interface {
		Execute(
			input *requestPresenter.CampaignProgressRequestPresenter,
			output *responsePresenter.CampaignProgressResponsePresenter,
		) (mvc.Result, error)
	}
)
