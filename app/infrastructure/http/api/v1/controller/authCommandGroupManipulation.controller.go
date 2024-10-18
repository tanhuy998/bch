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
		CreateCommandGroupUseCase                       usecasePort.IUseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse]
		AddUserToCommandGroupUseCase                    usecasePort.IUseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse]
		GetParitcipatedCommandGroupUseCase              usecasePort.IUseCase[requestPresenter.GetUserParticipatedCommandGroups, responsePresenter.GetUserParticipatedCommandGroups]
		GetTenantAllGroupsUseCase                       usecasePort.IUseCase[requestPresenter.GetTenantAllGroups, responsePresenter.GetTenantAllGroups]
		GetAssignmentUnAssignedCommandGroupUsersUseCase usecasePort.IUseCase[requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers, responsePresenter.GetAssignmentGroupUnAssignedCommandGroupUsers]
	}
)

func (this *AuthCommandGroupManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	activator.Handle(
		"GET", "/participated/user/{userUUID:uuid}", "GetParticipatedGroups",
		middleware.BindRequest[requestPresenter.GetUserParticipatedCommandGroups](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"GET", "/", "GetAllGroups",
		middleware.BindRequest[requestPresenter.GetTenantAllGroups](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"GET", "/users/unassigned/{assignmentGroupUUID:uuid}", "GetUnAssignedCommandGroupUsers",
		middleware.BindRequest[requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

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

func (this *AuthCommandGroupManipulationController) GetAllGroups(
	input *requestPresenter.GetTenantAllGroups,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetTenantAllGroupsUseCase.Execute(input),
	)
}

func (this *AuthCommandGroupManipulationController) GetParticipatedGroups(
	input *requestPresenter.GetUserParticipatedCommandGroups,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetParitcipatedCommandGroupUseCase.Execute(input),
	)
}

func (this *AuthCommandGroupManipulationController) GetUnAssignedCommandGroupUsers(
	input *requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetAssignmentUnAssignedCommandGroupUsersUseCase.Execute(input),
	)
}
