package assignmentServicePort

import (
	paginateServicePort "app/port/paginate"
	"context"

	"github.com/google/uuid"
)

type (
	IGetAssignmentGroups[Cursor_T comparable, Data_T any] interface {
		Serve(
			tenantUUID uuid.UUID, assignmentUUID uuid.UUID, paginator paginateServicePort.IPaginator[Cursor_T], ctx context.Context,
		) ([]Data_T, error)
	}
)
