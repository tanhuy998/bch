package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateTenant interface {
		Execute(
			input *requestPresenter.CreateTenantRequest,
			output *responsePresenter.CreateTenantResponse,
		) (mvc.Result, error)
	}
)
