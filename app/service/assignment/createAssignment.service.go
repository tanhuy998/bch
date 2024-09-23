package assignmentService

import (
	assignmentServicePort "app/adapter/assignment"
	"app/domain/model"
	libCommon "app/lib/common"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_DUPLICATE_ASSIGNMENT = errors.New("createAssignmentService error: dublicate assignment")
)

type (
	CreateAssignmentService struct {
		GetSingleAssignmentService assignmentServicePort.IGetSingleAssignnment
		AssignmentRepo             repository.IAssignment
	}
)

func (this *CreateAssignmentService) Serve(data *model.Assignment, ctx context.Context) (*model.Assignment, error) {

	existing, err := this.GetSingleAssignmentService.Search(data, ctx)

	if err != nil {

		return nil, err
	}

	if existing != nil {

		return nil, ERR_DUPLICATE_ASSIGNMENT
	}

	data.UUID = libCommon.PointerPrimitive(uuid.New())

	err = this.AssignmentRepo.Create(data, context.TODO())

	if err != nil {

		return nil, err
	}

	return data, nil
}
