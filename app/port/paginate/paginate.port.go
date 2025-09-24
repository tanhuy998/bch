package paginateServicePort

import (
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
		IAbstractPaginator
		GetPageNumber() int64
	}

	IAbstractPaginator interface {
		GetPageSize() uint64
	}

	ICursorPaginator[Cursor_T comparable] interface {
		IAbstractPaginator
		GetCursor() *Cursor_T
		IsPrevious() bool
		GetCursorDirection() CursorDirection
	}

	ICursorNillablePaginator[Cursor_T comparable] interface {
		CursorNilValue() *Cursor_T
	}

	// IFilterablePaginator interface {
	// 	ApplyPaginateFilter(filterGenerator repositoryAPI.IFilterGenerator)
	// }

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

	IGeneralPaginator interface {
		//IOffsetPaginator
		IPaginator[interface{}]
		GetCursor() interface{}
	}

	IGeneralPaginatorGetter interface {
		GetGeneralPaginator() IGeneralPaginator
	}

	IGeneralCursorPaginator interface {
		GetCursor() interface{}
		IsPrevious() bool
		GetCursorDirection() CursorDirection
	}
)

type (
	IGeneralPaginatorInput interface {
		GetGeneralPaginator() IGeneralPaginator
	}
)
