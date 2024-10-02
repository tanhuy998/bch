package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IAddUserToCommandGroup interface {
		Execute(
			input *requestPresenter.AddUserToCommandGroupRequest,
			output *responsePresenter.AddUserToCommandGroupResponse,
		) (mvc.Result, error)
	}
)
