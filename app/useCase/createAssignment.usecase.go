package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateAssignment interface {
		Execute(
			input *requestPresenter.CreateAssigmentRequest,
			output *responsePresenter.CreateAssignmentResponse,
		) (mvc.Result, error)
	}
)
