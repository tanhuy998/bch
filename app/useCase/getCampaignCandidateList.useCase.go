package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetCampaignCandidateList interface {
		Execute(
			*requestPresenter.GetCampaignCandidateListRequest,
			*responsePresenter.GetCampaignCandidateListResponse,
		) (mvc.Result, error)
	}
)
