package modifyAssignmentDomain

import (
	"app/internal/common"
	libError "app/internal/lib/error"
	"app/model"
	assignmentServicePort "app/port/assignment"
	"app/repository"
	repositoryAPI "app/repository/api"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type (
	ModifyAssignmentService struct {
		GetSingleAssignmentService assignmentServicePort.IGetSingleAssignnment
		AssignmentRepo             repository.IAssignment
	}
)

func (this *ModifyAssignmentService) Serve(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, dataModel *model.Assignment, ctx context.Context,
) error {

	switch existingAssignment, err := this.GetSingleAssignmentService.Serve(tenantUUID, assignmentUUID, ctx); {
	case err != nil:
		return err
	case existingAssignment == nil:
		return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment not found"))
	case existingAssignment.TenantUUID == nil:
		return libError.NewInternal(fmt.Errorf("wrong db data"))
	case *existingAssignment.TenantUUID != tenantUUID:
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment not in tenant"))
	}

	// similarQuery := bson.D{
	// 	{"title", dataModel.Title},
	// 	{"tenantUUID", tenantUUID},
	// 	{"assignmentUUID", assignmentUUID},
	// }

	// switch existingAssignment, err := this.AssignmentRepo.Find(similarQuery, ctx); {
	// case err != nil:
	// 	return err
	// case existingAssignment != nil:
	// 	return errors.Join(common.ERR_CONFLICT, fmt.Errorf("modified assignment conflict with existing assignment"))
	// }

	switch existingAssignment, err := this.AssignmentRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {
			filter.Field("tenantUUID").Equal(tenantUUID)
			filter.Field("title").Equal(dataModel.Title)
			filter.Field("assignmentUUID").Equal(assignmentUUID)
		},
	).FindOne(ctx); {
	case err != nil:
		return err
	case existingAssignment != nil:
		return errors.Join(common.ERR_CONFLICT, fmt.Errorf("modified assignment conflict with existing assignment"))
	}

	switch err := this.validateDeadLine(*dataModel.Deadline); {
	case err != nil:
		return err
	}

	dataModel.UUID = nil
	dataModel.TenantUUID = nil
	dataModel.CreatedBy = nil

	return this.AssignmentRepo.UpdateOneByUUID(assignmentUUID, dataModel, ctx)
}

func (this *ModifyAssignmentService) validateDeadLine(d time.Time) error {

	if time.Until(d) < 0 {

		return fmt.Errorf("modified assignment deadline must not be the time in the past")
	}

	return nil
}
