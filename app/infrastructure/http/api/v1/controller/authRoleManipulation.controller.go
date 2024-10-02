package controller

import (
	getAllRoleDomain "app/domain/auth/getAllRoles"
	grantCommandGroupRoleToUserDomain "app/domain/auth/grandCommandGroupRolesToUser"
	"app/infrastructure/http/common"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"
	libConfig "app/internal/lib/config"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthRoleManipulationController struct {
		*common.Controller
		GetAllRolesUseCase                  usecasePort.IUseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse]                                   // usecase.IGetAllRoles
		GrantCommandGroupRolesToUserUseCase usecasePort.IUseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse] // usecase.IGrantCommandGroupRolesToUser
	}
)

func (this *AuthRoleManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	this.bindDependencies(container)

	activator.Handle(
		"GET", "/", "GetAllRoles",
		middleware.BindRequest[requestPresenter.GetAllRolesRequest](
			container,
			middlewareHelper.UseAuthority,
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
		),
	)
}

func (this *AuthRoleManipulationController) bindDependencies(container *hero.Container) {

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse],
		getAllRoleDomain.GetAllRolesUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse],
		grantCommandGroupRoleToUserDomain.GrantCommandGroupRolesToUserUseCase,
	](container, nil)
}

func (this *AuthRoleManipulationController) GrantCommandGroupRolesToUser(
	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
	output *responsePresenter.GrantCommandGroupRolesToUserResponse,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GrantCommandGroupRolesToUserUseCase.Execute(input),
	)
}

func (this *AuthRoleManipulationController) GetAllRoles(
	input *requestPresenter.GetAllRolesRequest,
	output *responsePresenter.GetAllRolesResponse,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetAllRolesUseCase.Execute(input),
	)
}
