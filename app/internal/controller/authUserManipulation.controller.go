package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthUserManipulationController struct {
		CreateUserUsecase   usecase.ICreateUser
		GetGroupUserUsecase usecase.IGetGroupUsers
	}
)

func (this *AuthUserManipulationController) CreateUser(
	input *requestPresenter.CreateUserRequestPresenter,
	output *responsePresenter.CreateUserPresenter,
) (mvc.Result, error) {

	return this.CreateUserUsecase.Execute(input, output)
}

func (this *AuthUserManipulationController) GetGroupUsers(
	input *requestPresenter.GetGroupUsersRequest,
	output *responsePresenter.GetGroupUsersResponse,
) (mvc.Result, error) {

	return this.GetGroupUserUsecase.Execute(input, output)
}
