package assignmentServicePort

import (
	"app/model"
	"context"
	"errors"
)

var (
	ERR_DUPLICATE_ASSIGNMENT_GROUP = errors.New("createAssignmentGroup error: dublicate assignmentGroup")
	ERR_ASSIGNMENT_NOT_FOUND       = errors.New("createAssignmentGroup error: assignment not found")
)

type (
	ICreateAssignmentGroup interface {
		Serve(assignmentUUID string, ipnutData *model.AssignmentGroup, ctx context.Context) (*model.AssignmentGroup, error)
	}
)
