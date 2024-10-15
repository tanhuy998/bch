package assignmentServicePort

import (
	"app/model"
	"app/valueObject/requestInput"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetAssignmentPaginate interface {
		requestInput.IPaginationInput
		requestInput.IMongoCursorPaginationInput
		GetExpiredFilter() bool
	}

	GetAssignmentsFilter struct {
		PageNumber int
		Size       int
		Cursor     primitive.ObjectID
		Expired    bool
	}

	IGetAssignments interface {
		Serve(TenantUUID uuid.UUID, filter IGetAssignmentPaginate, ctx context.Context) ([]model.Assignment, error)
	}
)
