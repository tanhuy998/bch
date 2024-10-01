package createAssignmentGroupDomain

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	assignmentServicePort "app/src/port/assignment"
	"app/src/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	CreateAssignmentGroupService struct {
		AssignmentGroupRepo        repository.IAssignmentGroup
		GetSingleAssignmentService assignmentServicePort.IGetSingleAssignnment
	}
)

func (this *CreateAssignmentGroupService) Serve(
	assignmentUUID string, ipnutData *model.AssignmentGroup, ctx context.Context,
) (*model.AssignmentGroup, error) {

	assignment, err := this.GetSingleAssignmentService.Serve(assignmentUUID, ctx)

	if err != nil {

		return nil, err
	}

	if assignment == nil {

		return nil, assignmentServicePort.ERR_ASSIGNMENT_NOT_FOUND
	}

	existing, err := this.AssignmentGroupRepo.Find(
		bson.D{
			{"assignmentUUID", ipnutData.AssignmentUUID},
			{"tenantUUID", ipnutData.TenantUUID},
			{"commandGroupUUID", ipnutData.CommandGroupUUID},
			{"name", ipnutData.CommandGroupUUID},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	if existing != nil {

		return nil, assignmentServicePort.ERR_DUPLICATE_ASSIGNMENT_GROUP
	}

	ipnutData.UUID = libCommon.PointerPrimitive(uuid.New())

	err = this.AssignmentGroupRepo.Create(ipnutData, ctx)

	if err != nil {

		return nil, err
	}

	return ipnutData, nil
}
