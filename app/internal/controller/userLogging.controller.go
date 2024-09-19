package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	UserLoggingController struct {
		LoginUseCase usecase.ILogIn
	}
)

func (this *UserLoggingController) LogIn(
	input *requestPresenter.LoginRequest,
	output *responsePresenter.LoginResponse,
) (mvc.Result, error) {

	return this.LoginUseCase.Execute(input, output)
}
