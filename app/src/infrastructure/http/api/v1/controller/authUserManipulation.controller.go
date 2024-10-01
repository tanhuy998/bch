package controller

import (
	"app/src/infrastructure/http/middleware"
	"app/src/infrastructure/http/middleware/middlewareHelper"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	usecase "app/src/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthUserManipulationController struct {
		CreateUserUsecase   usecase.ICreateUser
		GetGroupUserUsecase usecase.IGetGroupUsers
		ModifyUserUsecase   usecase.IModifyUser
	}
)

func (this *AuthUserManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	activator.Handle(
		"POST", "/", "CreateUser",
		middleware.BindPresenters[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"GET", "/group/{groupUUID:uuid}", "GetGroupUsers",
		middleware.BindPresenters[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"PATCH", "/{userUUID:uuid}", "ModifyUser",
		middleware.BindPresenters[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)
}

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

func (this *AuthUserManipulationController) ModifyUser(
	input *requestPresenter.ModifyUserRequest,
	output *responsePresenter.ModifyUserResponse,
) (mvc.Result, error) {

	return this.ModifyUserUsecase.Execute(input, output)
}
