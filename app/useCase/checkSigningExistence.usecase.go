package usecase

import (
	requestPresenter "app/presenter/request"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICheckSigningExistence interface {
		Execute(input *requestPresenter.CheckSigningExistenceRequest) (mvc.Result, error)
	}
)
