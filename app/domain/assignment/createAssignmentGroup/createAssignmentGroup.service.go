package createAssignmentGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	"app/repository"
	"context"
	"errors"
	"fmt"

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
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, dataModel *model.AssignmentGroup, ctx context.Context,
) (*model.AssignmentGroup, error) {

	_, err := this.GetSingleAssignmentService.Serve(tenantUUID, assignmentUUID, ctx)

	if err != nil {

		return nil, err
	}

	existing, err := this.AssignmentGroupRepo.Find(
		bson.D{
			{"assignmentUUID", dataModel.AssignmentUUID},
			{"tenantUUID", dataModel.TenantUUID},
			{"name", dataModel.Name},
			// {"commandGroupUUID", dataModel.CommandGroupUUID},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	if existing != nil {

		return nil, errors.Join(common.ERR_BAD_REQUEST, fmt.Errorf("duplicate assignment group"))
	}

	dataModel.UUID = libCommon.PointerPrimitive(uuid.New())
	dataModel.AssignmentUUID = &assignmentUUID
	dataModel.TenantUUID = &tenantUUID

	err = this.AssignmentGroupRepo.Create(dataModel, ctx)

	if err != nil {

		return nil, err
	}

	return dataModel, nil
}
