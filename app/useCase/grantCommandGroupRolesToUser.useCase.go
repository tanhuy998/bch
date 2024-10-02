package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGrantCommandGroupRolesToUser interface {
		Execute(
			input *requestPresenter.GrantCommandGroupRolesToUserRequest,
			output *responsePresenter.GrantCommandGroupRolesToUserResponse,
		) (mvc.Result, error)
	}
)
