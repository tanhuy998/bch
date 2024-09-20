package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	UserLoggingController struct {
		LoginUseCase        usecase.ILogIn
		RefreshLoginUseCase usecase.IRefreshLogin
	}
)

func (this *UserLoggingController) LogIn(
	input *requestPresenter.LoginRequest,
	output *responsePresenter.LoginResponse,
) (mvc.Result, error) {

	return this.LoginUseCase.Execute(input, output)
}

func (this *UserLoggingController) Refresh(
	input *requestPresenter.RefreshLoginRequest,
	output *responsePresenter.RefreshLoginResponse,
) (mvc.Result, error) {

	return this.RefreshLoginUseCase.Execute(input, output)
}

func (this *UserLoggingController) Temp() {

}
