package usecase

import (
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetParticipatedCommandGroups interface {
		Execute(
			input *requestPresenter.ReportParticipatedGroups,
			output *responsePresenter.ReportParticipatedGroups,
		) (mvc.Result, error)
	}
)
