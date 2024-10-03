package boundedContext

import (
	libConfig "app/internal/lib/config"
	assignmentServicePort "app/port/assignment"

	createAssignmentDomain "app/domain/assignment/createAssignment"
	createAssignmentGroupDomain "app/domain/assignment/createAssignmentGroup"
	getSingleAssignmentDomain "app/domain/assignment/getSingleAssignment"
	getSingleAssignmentGroupDomain "app/domain/assignment/getSingleAssignmentGroup"

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

	libConfig.BindDependency[assignmentServicePort.IGetSingleAssignnment, getSingleAssignmentDomain.GetSingleAssignmentService](container, nil)
	libConfig.BindDependency[assignmentServicePort.IGetSingleAssignmentGroup, getSingleAssignmentGroupDomain.GetSingleAssignmentGroupService](container, nil)
	libConfig.BindDependency[assignmentServicePort.ICreateAssignment, createAssignmentDomain.CreateAssignmentService](container, nil)
	libConfig.BindDependency[assignmentServicePort.ICreateAssignmentGroup, createAssignmentGroupDomain.CreateAssignmentGroupService](container, nil)

	container.Register(new(AssignmentBoundedContext)).Explicitly().EnableStructDependents()
}
