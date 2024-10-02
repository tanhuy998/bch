package createAssignmentGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	"app/repository"
	"context"
	"errors"

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

		return nil, errors.Join(common.ERR_BAD_REQUEST, errors.New("assignment not found")) // assignmentServicePort.ERR_ASSIGNMENT_NOT_FOUND
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

		return nil, errors.Join(common.ERR_BAD_REQUEST, errors.New("duplicate assignment group")) //assignmentServicePort.ERR_DUPLICATE_ASSIGNMENT_GROUP
	}

	ipnutData.UUID = libCommon.PointerPrimitive(uuid.New())

	err = this.AssignmentGroupRepo.Create(ipnutData, ctx)

	if err != nil {

		return nil, errors.Join(common.ERR_INTERNAL, err)
	}

	return ipnutData, nil
}
