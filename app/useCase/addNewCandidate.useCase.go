package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IAddNewCandidate interface {
		Execute(
			input *requestPresenter.AddCandidateRequest,
			output *responsePresenter.AddNewCandidateResponse,
		) (mvc.Result, error)
	}
)
