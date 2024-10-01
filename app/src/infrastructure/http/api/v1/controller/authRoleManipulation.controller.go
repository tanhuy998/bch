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
	AuthRoleManipulationController struct {
		GetAllRolesUseCase                  usecase.IGetAllRoles
		GrantCommandGroupRolesToUserUseCase usecase.IGrantCommandGroupRolesToUser
	}
)

func (this *AuthRoleManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	activator.Handle(
		"GET", "/", "GetAllRoles",
		middleware.BindPresenters[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse](
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
		middleware.BindPresenters[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)
}

func (this *AuthRoleManipulationController) GrantCommandGroupRolesToUser(
	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
	output *responsePresenter.GrantCommandGroupRolesToUserResponse,
) (mvc.Result, error) {

	return this.GrantCommandGroupRolesToUserUseCase.Execute(input, output)
}

func (this *AuthRoleManipulationController) GetAllRoles(
	input *requestPresenter.GetAllRolesRequest,
	output *responsePresenter.GetAllRolesResponse,
) (mvc.Result, error) {

	return this.GetAllRolesUseCase.Execute(input, output)
}
