package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IModifyExistingCandidate interface {
		Execute(
			input *requestPresenter.ModifyExistingCandidateRequest,
			output *responsePresenter.ModifyExistingCandidateResponse,
		) (mvc.Result, error)
	}
)
