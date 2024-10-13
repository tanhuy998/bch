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
	AssignmentController struct {
		common.Controller
		CreateAssignmentUseCase                     usecasePort.IUseCase[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse]
		GetSingleAssignmentUseCase                  usecasePort.IUseCase[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse]
		CreateAssignmentGroupUseCase                usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse]
		ModifyAssignmentUseCase                     usecasePort.IUseCase[requestPresenter.ModifyAssignment, responsePresenter.ModifyAssignment]
		AddCommandGroupUserToAssignmentGroupUseCase usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupMember, responsePresenter.CreateAssignmentGroupMemeber]
	}
)

func (this *AssignmentController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Dependencies()

	// default
	activator.Router().Use(
		middleware.Auth(
			container,
		),
	)

	activator.Handle(
		"GET", "/{uuid:uuid}", "GetSingleAssignment",
		middleware.BindRequest[requestPresenter.GetSingleAssignmentRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"POST", "/", "CreateAssignment",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
		middleware.BindRequest[requestPresenter.CreateAssigmentRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"POST", "/{assignmentUUID:uuid}/group/command/{commandGroupUUID:uuid}", "CreateAssignmentGroup",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
		middleware.BindRequest[requestPresenter.CreateAssignmentGroupRequest](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"PATCH", "/{assignmentUUID:uuid}", "ModifyAssignment",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
		middleware.BindRequest[requestPresenter.ModifyAssignment](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)

	activator.Handle(
		"POST", "/group/{groupUUID:uuid}/member", "CreateAssignmentGroupMember",
		middleware.Auth(
			container,
			middlewareHelper.AuthRequiredTenantAgentExceptMeetRoles("COMMANDER"),
		),
		middleware.BindRequest[requestPresenter.CreateAssignmentGroupMember](
			container,
			middlewareHelper.UseAuthority,
			middlewareHelper.UseTenantMapping,
		),
	)
}

func (this *AssignmentController) BindDependencies(container *hero.Container) common.IController {

	return this
}

func (this *AssignmentController) GetSingleAssignment(
	input *requestPresenter.GetSingleAssignmentRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.GetSingleAssignmentUseCase.Execute(input),
	)
}

func (this *AssignmentController) CreateAssignment(
	input *requestPresenter.CreateAssigmentRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CreateAssignmentUseCase.Execute(input),
	)
}

func (this *AssignmentController) CreateAssignmentGroup(
	input *requestPresenter.CreateAssignmentGroupRequest,
) (mvc.Result, error) {

	return this.ResultOf(
		this.CreateAssignmentGroupUseCase.Execute(input),
	)
}

func (this *AssignmentController) ModifyAssignment(
	input *requestPresenter.ModifyAssignment,
) (mvc.Result, error) {

	return this.ResultOf(
		this.ModifyAssignmentUseCase.Execute(input),
	)
}

func (this *AssignmentController) CreateAssignmentGroupMember(
	input *requestPresenter.CreateAssignmentGroupMember,
) (mvc.Result, error) {

	return this.ResultOf(
		this.AddCommandGroupUserToAssignmentGroupUseCase.Execute(input),
	)
}
