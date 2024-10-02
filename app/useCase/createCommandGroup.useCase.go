package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateCommandGroup interface {
		Execute(
			input *requestPresenter.CreateCommandGroupRequest,
			output *responsePresenter.CreateCommandGroupResponse,
		) (mvc.Result, error)
	}
)
