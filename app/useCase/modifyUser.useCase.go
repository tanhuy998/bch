package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IModifyUser interface {
		Execute(
			input *requestPresenter.ModifyUserRequest,
			output *responsePresenter.ModifyUserResponse,
		) (mvc.Result, error)
	}
)
