package assignmentServicePort

import (
	"app/model"
	paginateServicePort "app/port/paginate"
	"context"

	"github.com/google/uuid"
)

type (
	IGetAssignmentGroups[Cursor_T comparable] interface {
		Serve(
			tenantUUID uuid.UUID, assignmentUUID uuid.UUID, paginator paginateServicePort.IPaginator[Cursor_T], ctx context.Context,
		) ([]model.AssignmentGroup, error)
	}
)
