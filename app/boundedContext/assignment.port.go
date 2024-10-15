package boundedContext

import (
	libConfig "app/internal/lib/config"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	createAssignmentDomain "app/domain/assignment/createAssignment"
	createAssignmentGroupDomain "app/domain/assignment/createAssignmentGroup"
	"app/domain/assignment/createAssignmentGroupMemberDomain"
	getAssignmentsDomain "app/domain/assignment/getAssignments"
	getSingleAssignmentDomain "app/domain/assignment/getSingleAssignment"
	getSingleAssignmentGroupDomain "app/domain/assignment/getSingleAssignmentGroup"
	modifyAssignmentDomain "app/domain/assignment/modifyAssignment"

	"github.com/kataras/iris/v12/hero"
)

type (
	AssignmentBoundedContext struct {
		assignmentServicePort.ICreateAssignment
		assignmentServicePort.ICreateAssignmentGroup
		assignmentServicePort.IGetSingleAssignnment
		assignmentServicePort.IGetSingleAssignmentGroup
	}
)

func RegisterAssignmentBoundedContext(container *hero.Container) {

	libConfig.BindDependency[assignmentServicePort.IGetAssignments, getAssignmentsDomain.GetAssignmentsService](container, nil)
	libConfig.BindDependency[assignmentServicePort.IGetSingleAssignnment, getSingleAssignmentDomain.GetSingleAssignmentService](container, nil)
	libConfig.BindDependency[assignmentServicePort.IGetSingleAssignmentGroup, getSingleAssignmentGroupDomain.GetSingleAssignmentGroupService](container, nil)
	libConfig.BindDependency[assignmentServicePort.ICreateAssignment, createAssignmentDomain.CreateAssignmentService](container, nil)
	libConfig.BindDependency[assignmentServicePort.ICreateAssignmentGroup, createAssignmentGroupDomain.CreateAssignmentGroupService](container, nil)
	libConfig.BindDependency[assignmentServicePort.IModifyAssignment, modifyAssignmentDomain.ModifyAssignmentService](container, nil)

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
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.ModifyAssignment, responsePresenter.ModifyAssignment],
		modifyAssignmentDomain.ModifyAssignmentUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupMember, responsePresenter.CreateAssignmentGroupMemeber],
		createAssignmentGroupMemberDomain.CreateAssignmentGroupMemberUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAssignments, responsePresenter.GetAssignments],
		getAssignmentsDomain.GetAssignmentUseCase,
	](container, nil)

	container.Register(new(AssignmentBoundedContext)).Explicitly().EnableStructDependents()
}
