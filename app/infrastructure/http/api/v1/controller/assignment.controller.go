package controller

import (
	createAssignmentDomain "app/domain/assignment/createAssignment"
	createAssignmentGroupDomain "app/domain/assignment/createAssignmentGroup"
	getSingleAssignmentDomain "app/domain/assignment/getSingleAssignment"
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
	AssignmentController struct {
		*common.Controller
		CreateAssignmentUseCase      usecasePort.IUseCase[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse]
		GetSingleAssignmentUseCase   usecasePort.IUseCase[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse]
		CreateAssignmentGroupUseCase usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse]
	}
)

func (this *AssignmentController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	this.bindDependencies(container)

	activator.Handle(
		"GET", "/{uuid:uuid}", "GetSingleAssignment",
		middleware.BindRequest[requestPresenter.GetSingleAssignmentRequest](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"POST", "/", "CreateAssignment",
		middleware.BindRequest[requestPresenter.CreateAssigmentRequest](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"POST", "/{assignmentUUID:uuid}/group", "CreateAssignmentGroup",
		middleware.BindRequest[requestPresenter.CreateAssignmentGroupRequest](
			container,
			middlewareHelper.UseAuthority,
		),
	)
}

func (this *AssignmentController) bindDependencies(container *hero.Container) {

	// libConfig.BindDependency[assignmentServicePort.IGetSingleAssignnment, getSingleAssignmentDomain.GetSingleAssignmentService](container, nil)
	// libConfig.BindDependency[assignmentServicePort.ICreateAssignment, createAssignmentDomain.CreateAssignmentService](container, nil)
	// libConfig.BindDependency[assignmentServicePort.ICreateAssignmentGroup, createAssignmentGroupDomain.CreateAssignmentGroupService](container, nil)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse],
		createAssignmentDomain.CreateAssignmentUseCase,
	](container, nil)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse],
		getSingleAssignmentDomain.GetSingleAssignmentUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse],
		createAssignmentGroupDomain.CreateAssignmentGroupUseCase,
	](container, nil)
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
