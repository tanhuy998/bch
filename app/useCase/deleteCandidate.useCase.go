package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IDeleteCandidate interface {
		Execute(
			input *requestPresenter.DeleteCandidateRequest,
			output *responsePresenter.DeleteCandidateResponse,
		) (mvc.Result, error)
	}
)
