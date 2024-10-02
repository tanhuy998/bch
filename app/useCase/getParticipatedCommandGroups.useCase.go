package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetParticipatedCommandGroups interface {
		Execute(
			input *requestPresenter.GetParticipatedGroups,
			output *responsePresenter.GetParticipatedGroups,
		) (mvc.Result, error)
	}
)
