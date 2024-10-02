package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateUser interface {
		Execute(
			input *requestPresenter.CreateUserRequestPresenter,
			output *responsePresenter.CreateUserPresenter,
		) (mvc.Result, error)
	}
)
