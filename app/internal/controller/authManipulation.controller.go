package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthManipulationController struct {
		CreateUserUsecase usecase.ICreateUser
	}
)

func (this *AuthManipulationController) CreateUser(
	input *requestPresenter.CreateUserRequestPresenter,
	output *responsePresenter.CreateUserPresenter,
) (mvc.Result, error) {

	return this.CreateUserUsecase.Execute(input, output)
}
