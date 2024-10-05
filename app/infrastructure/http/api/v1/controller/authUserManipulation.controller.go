package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthUserManipulationController struct {
		common.Controller
		CreateUserUsecase   usecasePort.IUseCase[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter] // usecase.ICreateUser
		GetGroupUserUsecase usecasePort.IUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse]     // usecase.IGetGroupUsers
		ModifyUserUsecase   usecasePort.IUseCase[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse]           // usecase.IModifyUser
	}
)

func (this *AuthUserManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

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

func (this *AuthUserManipulationController) BindDependencies(container *hero.Container) common.IController {

	return this
}

func (this *AuthUserManipulationController) CreateUser(
	input *requestPresenter.CreateUserRequestPresenter,
	output *responsePresenter.CreateUserPresenter,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CreateUserUsecase.Execute(input),
	)
}

func (this *AuthUserManipulationController) GetGroupUsers(
	input *requestPresenter.GetGroupUsersRequest,
	output *responsePresenter.GetGroupUsersResponse,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetGroupUserUsecase.Execute(input),
	)
}

func (this *AuthUserManipulationController) ModifyUser(
	input *requestPresenter.ModifyUserRequest,
	output *responsePresenter.ModifyUserResponse,
) (mvc.Result, error) {

	return this.ResultOf(
		this.ModifyUserUsecase.Execute(input),
	)
}
