package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateAssignmentGroup interface {
		Execute(
			input *requestPresenter.CreateAssignmentGroupRequest,
			output *responsePresenter.CreateAssignmentGroupResponse,
		) (mvc.Result, error)
	}
)
