package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetGroupUsers interface {
		Execute(
			input *requestPresenter.GetGroupUsersRequest,
			output *responsePresenter.GetGroupUsersResponse,
		) (mvc.Result, error)
	}
)
