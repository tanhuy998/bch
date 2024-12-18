package paginateServicePort

import (
	repositoryAPI "app/repository/api"
	"context"

	"github.com/google/uuid"
)

type (
	CursorDirection uint
)

const (
	CURSOR_DIRECTION_PREVIOUS CursorDirection = iota
	CURSOR_DIRECTION_NEXT
)

type (
	IPaginateService[Entity_T any, Cursor_T comparable] interface {
		Serve(
			//tenantUUID uuid.UUID, page uint64, size uint64, cursor *Cursor_T, isPrev bool, ctx context.Context,
			tenantUUID uuid.UUID, paginator IPaginator[Cursor_T], ctx context.Context,
		) ([]Entity_T, error)
	}

	IOffsetPaginator interface {
		GetPageNumber() uint64
		GetPageSize() uint64
	}

	ICursorPaginator[Cursor_T comparable] interface {
		GetCursor() *Cursor_T
		IsPrevious() bool
		GetCursorDirection() CursorDirection
	}

	ICursorNillablePaginator[Cursor_T comparable] interface {
		CursorNilValue() *Cursor_T
	}

	IFilterablePaginator interface {
		ApplyPaginateFilter(filterGenerator repositoryAPI.IFilterGenerator)
	}

	IPaginateProjector interface {
		Select(fields ...string)
		ExcludeField(fields ...string)
	}

	IProjectionPaginator interface {
		ApplyPaginateProjection(projector IPaginateProjector)
	}

	IPaginator[Cursor_T comparable] interface {
		IOffsetPaginator
		//ICursorPaginator[Cursor_T]
	}
)
