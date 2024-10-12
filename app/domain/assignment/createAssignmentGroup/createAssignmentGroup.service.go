package createAssignmentGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	assignmentServicePort "app/port/assignment"
	authServicePort "app/port/auth"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	CreateAssignmentGroupService struct {
		AssignmentGroupRepo          repository.IAssignmentGroup
		GetSingleAssignmentService   assignmentServicePort.IGetSingleAssignnment
		GetSingleCommandGroupService authServicePort.IGetSingleCommandGroup
	}
)

func (this *CreateAssignmentGroupService) Serve(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, commandGroupUUID uuid.UUID, dataModel *model.AssignmentGroup, ctx context.Context,
) (*model.AssignmentGroup, error) {

	switch _, err := this.GetSingleAssignmentService.Serve(tenantUUID, assignmentUUID, ctx); {
	case err != nil:
		return nil, err
	}

	switch existingCommandGroup, err := this.GetSingleCommandGroupService.Serve(commandGroupUUID, ctx); {
	case err != nil:
		return nil, err
	case existingCommandGroup == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("command group not found"))
	case *existingCommandGroup.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("commond group not in tenant"))
	}

	similarQuery := bson.D{
		{"assignmentUUID", assignmentUUID},
		{"tenantUUID", tenantUUID},
		{"name", dataModel.Name},
	}

	switch existing, err := this.AssignmentGroupRepo.Find(similarQuery, ctx); {
	case err != nil:
		return nil, err
	case existing != nil:
		return nil, errors.Join(common.ERR_CONFLICT, fmt.Errorf("there is similar assignment group in the current tenant assignment, try another payload"))
	}

	dataModel.UUID = libCommon.PointerPrimitive(uuid.New())
	dataModel.AssignmentUUID = &assignmentUUID
	dataModel.CommandGroupUUID = &commandGroupUUID
	dataModel.TenantUUID = &tenantUUID

	err := this.AssignmentGroupRepo.Create(dataModel, ctx)

	if err != nil {

		return nil, err
	}

	return dataModel, nil
}
