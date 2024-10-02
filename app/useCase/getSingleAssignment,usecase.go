package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleAssignment interface {
		Execute(
			input *requestPresenter.GetSingleAssignmentRequest,
			output *responsePresenter.GetSingleAssignmentResponse,
		) (mvc.Result, error)
	}
)
