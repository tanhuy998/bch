package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"
	"app/model"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthUserManipulationController struct {
		common.Controller
		CreateUserUsecase     usecasePort.IUseCase[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter] // usecase.ICreateUser
		GetGroupUserUsecase   usecasePort.IUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse]     // usecase.IGetGroupUsers
		ModifyUserUsecase     usecasePort.IUseCase[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse]           // usecase.IModifyUser
		GetTenantUsersUseCase usecasePort.IUseCase[requestPresenter.GetTenantUsers, responsePresenter.GetTenantUsers[model.User]]
	}
)

func (this *AuthUserManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	activator.Router().Use(
		middleware.Auth(
			container,
		),
	)

	activator.Handle(
		"POST", "/", "CreateUser",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
		middleware.BindRequest[requestPresenter.CreateUserRequestPresenter](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"GET", "/group/{groupUUID:uuid}", "GetGroupUsers",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequiredTenantAgentExceptMeetRoles("COMMANDER"),
		),
		middleware.BindRequest[requestPresenter.GetGroupUsersRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"GET", "/", "GetTenantUsers",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
		middleware.BindPresenters[requestPresenter.GetTenantUsers, responsePresenter.GetTenantUsers[model.User]](
			container,
			middlewareHelper.UseTenantMapping,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"PATCH", "/{userUUID:uuid}", "ModifyUser",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
		middleware.BindRequest[requestPresenter.ModifyUserRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"GET", "/group/{groupUUID:uuid}", "GetGroupUsers",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequiredTenantAgentExceptMeetRoles("COMMANDER"),
		),
		middleware.BindRequest[requestPresenter.GetGroupUsersRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)
}

func (this *AuthUserManipulationController) BindDependencies(container *hero.Container) common.IController {

	return this
}

func (this *AuthUserManipulationController) CreateUser(
	input *requestPresenter.CreateUserRequestPresenter,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CreateUserUsecase.Execute(input),
	)
}

func (this *AuthUserManipulationController) GetGroupUsers(
	input *requestPresenter.GetGroupUsersRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetGroupUserUsecase.Execute(input),
	)
}

func (this *AuthUserManipulationController) ModifyUser(
	input *requestPresenter.ModifyUserRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.ModifyUserUsecase.Execute(input),
	)
}

func (this *AuthUserManipulationController) GetTenantUsers(
	input *requestPresenter.GetTenantUsers,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetTenantUsersUseCase.Execute(input),
	)
}
