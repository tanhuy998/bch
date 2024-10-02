package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetCampaignUnSignedCandidates interface {
		Execute(
			input *requestPresenter.GetCampaignUnSignedCandidates,
			output *responsePresenter.GetCampaignUnSignedCandidates,
		) (mvc.Result, error)
	}
)
