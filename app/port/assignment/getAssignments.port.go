package assignmentServicePort

import (
	"app/model"
	paginateServicePort "app/port/paginate"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetAssignmentPaginate[Cursor_T comparable] interface {
		// requestInput.IPaginationInput
		// requestInput.IMongoCursorPaginationInput
		paginateServicePort.IPaginator[Cursor_T]
		GetExpiredFilter() bool
	}

	GetAssignmentsFilter struct {
		PageNumber int
		Size       int
		Cursor     primitive.ObjectID
		Expired    bool
	}

	IGetAssignments[Cursor_T comparable] interface {
		Serve(TenantUUID uuid.UUID, filter IGetAssignmentPaginate[Cursor_T], ctx context.Context) ([]model.Assignment, error)
	}
)
