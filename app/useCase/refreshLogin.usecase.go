package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"

	"github.com/kataras/iris/v12/mvc"
)

var (
	ERR_REFRESH_NO_CONTEXT = errors.New("refresh login usecase: no context")
)

type (
	IRefreshLogin interface {
		Execute(
			input *requestPresenter.RefreshLoginRequest,
			output *responsePresenter.RefreshLoginResponse,
		) (mvc.Result, error)
	}
)
