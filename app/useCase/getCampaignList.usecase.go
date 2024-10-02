package usecase

import (
	"app/model"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"

	"github.com/kataras/iris/v12/mvc"
)

type (
	RetrievedData_T = repository.PaginationPack[model.Campaign]

	IGetCampaignList interface {
		Execute(
			input *requestPresenter.GetCampaignListRequest,
			output *responsePresenter.GetCampaignListResponse,
		) (mvc.Result, error)
	}
)
