package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleCandidateSigningInfo interface {
		Execute(
			input *requestPresenter.GetSingleCandidateSigningInfoRequest,
			output *responsePresenter.GetSingleCandidateSigningInfoResponse,
		) (mvc.Result, error)
	}
)
