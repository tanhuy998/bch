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
	AuthCommandGroupManipulationController struct {
		common.Controller
		CreateCommandGroupUseCase          usecasePort.IUseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse]
		AddUserToCommandGroupUseCase       usecasePort.IUseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse]
		GetParitcipatedCommandGroupUseCase usecasePort.IUseCase[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups]
	}
)

func (this *AuthCommandGroupManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	activator.Handle(
		"GET", "/participated/user/{userUUID:uuid}", "GetParticipatedGroups",
		middleware.BindRequest[requestPresenter.GetParticipatedGroups](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	// activator.Handle(
	// 	"GET", "/", "GetAllGroups",
	// )

	activator.Handle(
		"POST", "/", "CreateGroup",
		middleware.BindRequest[requestPresenter.CreateCommandGroupRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"POST", "/{groupUUID:uuid}/user/{userUUID:uuid}", "AddUserToGroup",
		middleware.BindRequest[requestPresenter.AddUserToCommandGroupRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)
}

func (this *AuthCommandGroupManipulationController) BindDependencies(container *hero.Container) common.IController {

	return this
}

func (this *AuthCommandGroupManipulationController) CreateGroup(
	input *requestPresenter.CreateCommandGroupRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CreateCommandGroupUseCase.Execute(input),
	)
}

func (this *AuthCommandGroupManipulationController) AddUserToGroup(
	input *requestPresenter.AddUserToCommandGroupRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.AddUserToCommandGroupUseCase.Execute(input),
	)
}

func (this *AuthCommandGroupManipulationController) GetAllGroups() {

}

func (this *AuthCommandGroupManipulationController) GetParticipatedGroups(
	input *requestPresenter.GetParticipatedGroups,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetParitcipatedCommandGroupUseCase.Execute(input),
	)
}

// func (this *AuthCommandGroupManipulationController) GrantCommandGroupRolesToUser(
// 	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
// 	output *responsePresenter.GrantCommandGroupRolesToUserResponse,
// ) (mvc.Result, error) {

// 	return this.GrantCommandGroupRolesToUserUseCase.Execute(input, output)
// }
