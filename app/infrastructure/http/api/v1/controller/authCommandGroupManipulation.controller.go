package controller

import (
	addUserToCommandGroupDomain "app/domain/auth/addUserToCommandGroup"
	createCommandGroupDomain "app/domain/auth/createCommandGroup"
	getUserParticipatedCommandGroupDomain "app/domain/auth/getUserParticipatedGroups"
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
	AuthCommandGroupManipulationController struct {
		*common.Controller
		CreateCommandGroupUseCase          usecasePort.IUseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse]
		AddUserToCommandGroupUseCase       usecasePort.IUseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse]
		GetParitcipatedCommandGroupUseCase usecasePort.IUseCase[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups]
	}
)

func (this *AuthCommandGroupManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	this.bindDependencies(container)

	activator.Handle(
		"GET", "/participated/user/{userUUID:uuid}", "GetParticipatedGroups",
		middleware.BindRequest[requestPresenter.GetParticipatedGroups](
			container,
			middlewareHelper.UseAuthority,
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
		),
	)

	activator.Handle(
		"POST", "/{groupUUID:uuid}/user/{userUUID:uuid}", "AddUserToGroup",
		middleware.BindRequest[requestPresenter.AddUserToCommandGroupRequest](
			container,
			middlewareHelper.UseAuthority,
		),
	)
}

func (this *AuthCommandGroupManipulationController) bindDependencies(container *hero.Container) {

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse],
		createCommandGroupDomain.CreateCommandGroupUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse],
		addUserToCommandGroupDomain.AddUserToCommandGroupUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups],
		getUserParticipatedCommandGroupDomain.GetParticipatedCommandGroupsUseCase,
	](container, nil)
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
	output *responsePresenter.AddUserToCommandGroupResponse,
) (mvc.Result, error) {

	return this.ResultOf(
		this.AddUserToCommandGroupUseCase.Execute(input),
	)
}

func (this *AuthCommandGroupManipulationController) GetAllGroups() {

}

func (this *AuthCommandGroupManipulationController) GetParticipatedGroups(
	input *requestPresenter.GetParticipatedGroups,
	output *responsePresenter.GetParticipatedGroups,
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
