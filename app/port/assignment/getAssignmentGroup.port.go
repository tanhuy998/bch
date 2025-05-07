package assignmentServicePort

import (
	paginateServicePort "app/port/paginate"
	"app/unitOfWork/aggregate"
	"context"

	"github.com/google/uuid"
)

type (
	IGetAssignmentInputContext[Cursor_T comparable] interface {
		context.Context
		aggregate.IDomainContext
		paginateServicePort.IPaginateContext[Cursor_T]
		GetRequestedAssignmentUUID() uuid.UUID
	}

	// IGetAssignmentGroups[Cursor_T comparable, Data_T any] interface {
	// 	Serve(
	// 		tenantUUID uuid.UUID, assignmentUUID uuid.UUID, paginator paginateServicePort.IPaginator[Cursor_T], ctx context.Context,
	// 	) ([]Data_T, error)
	// }

	IGetAssignmentGroups[Cursor_T comparable, Data_T any] interface {
		Serve(
			inputCtx IGetAssignmentInputContext[Cursor_T],
		) ([]Data_T, error)
	}
)
