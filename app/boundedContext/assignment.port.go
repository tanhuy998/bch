package boundedContext

import (
	"app/domain"
	irisIoc "app/internal/lib/iris/ioc"
	"app/model"
	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	createAssignmentDomain "app/domain/assignment/createAssignment"
	createAssignmentGroupDomain "app/domain/assignment/createAssignmentGroup"
	"app/domain/assignment/createAssignmentGroupMemberDomain"
	"app/domain/assignment/getAssignmentGroupsDomain"
	getAssignmentsDomain "app/domain/assignment/getAssignments"
	getSingleAssignmentDomain "app/domain/assignment/getSingleAssignment"
	getSingleAssignmentGroupDomain "app/domain/assignment/getSingleAssignmentGroup"
	checkAssignmentParticipantDomain "app/domain/auth/checkAssignmenParticipant"

	modifyAssignmentDomain "app/domain/assignment/modifyAssignment"

	"github.com/kataras/iris/v12/hero"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	irisIoc.BindDependency[assignmentServicePort.IGetAssignments[primitive.ObjectID], getAssignmentsDomain.GetAssignmentsService](container, nil)
	irisIoc.BindDependency[assignmentServicePort.IGetAssignmentGroups[domain.PaginateCursorType, model.AssignmentGroup], getAssignmentGroupsDomain.GetAssignmentGroupsService](container, nil)
	irisIoc.BindDependency[assignmentServicePort.IGetSingleAssignnment, getSingleAssignmentDomain.GetSingleAssignmentService](container, nil)
	irisIoc.BindDependency[assignmentServicePort.IGetSingleAssignmentGroup, getSingleAssignmentGroupDomain.GetSingleAssignmentGroupService](container, nil)
	irisIoc.BindDependency[assignmentServicePort.ICreateAssignment, createAssignmentDomain.CreateAssignmentService](container, nil)
	irisIoc.BindDependency[assignmentServicePort.ICreateAssignmentGroup, createAssignmentGroupDomain.CreateAssignmentGroupService](container, nil)
	irisIoc.BindDependency[assignmentServicePort.ICreateAssignmentGroupMember, createAssignmentGroupMemberDomain.CreateAssignmentGroupMemberService](container, nil)
	irisIoc.BindDependency[assignmentServicePort.IModifyAssignment, modifyAssignmentDomain.ModifyAssignmentService](container, nil)

	irisIoc.BindDependency[assignmentServicePort.ICheckAssigmnetParticipant, checkAssignmentParticipantDomain.CheckAssignmentPariticipantService](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse],
		createAssignmentDomain.CreateAssignmentUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse],
		getSingleAssignmentDomain.GetSingleAssignmentUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse],
		createAssignmentGroupDomain.CreateAssignmentGroupUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.ModifyAssignment, responsePresenter.ModifyAssignment],
		modifyAssignmentDomain.ModifyAssignmentUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateAssignmentGroupMember, responsePresenter.CreateAssignmentGroupMemeber],
		createAssignmentGroupMemberDomain.CreateAssignmentGroupMemberUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAssignments, responsePresenter.GetAssignments],
		getAssignmentsDomain.GetAssignmentUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAssignmentGroups, responsePresenter.GetAssignmentGroups[model.AssignmentGroup]],
		getAssignmentGroupsDomain.GetAssignmentGroupsUseCase,
	](container, nil)

	container.Register(new(AssignmentBoundedContext)).Explicitly().EnableStructDependents()
}
