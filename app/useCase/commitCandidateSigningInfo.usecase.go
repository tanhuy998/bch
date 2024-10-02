package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICommitCandidateSigningInfo interface {
		Execute(
			input *requestPresenter.CommitCandidateSigningInfoRequest,
			outout *responsePresenter.CommitCandidateSigningInfoResponse,
		) (mvc.Result, error)
	}
)
