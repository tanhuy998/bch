package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleCandidateByUUID interface {
		Execute(
			input *requestPresenter.GetSingleCandidateRequest,
			output *responsePresenter.GetSingleCandidateResponse,
		) (mvc.Result, error)
	}
)
