package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICommitSpecificSigningInfo interface {
		Execute(
			input *requestPresenter.CommitSpecificSigningInfo,
			output *responsePresenter.CommitSpecificSigningInfoResponse,
		) (mvc.Result, error)
	}
)
