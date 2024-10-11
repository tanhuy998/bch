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
	AuthRoleManipulationController struct {
		common.Controller
		GetAllRolesUseCase                  usecasePort.IUseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse]                                   // usecase.IGetAllRoles
		GrantCommandGroupRolesToUserUseCase usecasePort.IUseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse] // usecase.IGrantCommandGroupRolesToUser
	}
)

func (this *AuthRoleManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	activator.Handle(
		"GET", "/", "GetAllRoles",
		middleware.BindRequest[requestPresenter.GetAllRolesRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	// activator.Handle(
	// 	"POST", "/", "CreateRole",
	// 	middleware.BindPresenters[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse](container),
	// )

	activator.Handle(
		"POST", "/group/{groupUUID}/user/{userUUID}", "GrantCommandGroupRolesToUser",
		middleware.BindRequest[requestPresenter.GrantCommandGroupRolesToUserRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)
}

func (this *AuthRoleManipulationController) BindDependencies(container *hero.Container) common.IController {

	return this
}

func (this *AuthRoleManipulationController) GrantCommandGroupRolesToUser(
	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GrantCommandGroupRolesToUserUseCase.Execute(input),
	)
}

func (this *AuthRoleManipulationController) GetAllRoles(
	input *requestPresenter.GetAllRolesRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetAllRolesUseCase.Execute(input),
	)
}
