package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetCampaignSignedCandidates interface {
		Execute(
			input *requestPresenter.GetCampaignSignedCandidatesRequest,
			output *responsePresenter.GetCampaignSignedCandidatesResponse,
		) (mvc.Result, error)
	}
)
