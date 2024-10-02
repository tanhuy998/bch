package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IDeleteCampaign interface {
		Execute(
			*requestPresenter.DeleteCampaignRequest,
			*responsePresenter.DeleteCampaignResponse,
		) (mvc.Result, error)
	}
)
