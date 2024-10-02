package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"

	"github.com/kataras/iris/v12/mvc"
)

var (
	ERR_NIL_CONTEXT = errors.New("login usecase error: nil context")
)

type (
	ILogIn interface {
		Execute(*requestPresenter.LoginRequest, *responsePresenter.LoginResponse) (mvc.Result, error)
	}
)
