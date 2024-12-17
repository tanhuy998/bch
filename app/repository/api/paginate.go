package repositoryAPI

import (
	"context"
)

type (
	ICursorPaginationRepository[Model_T any] interface {
		FindNext(
			cursorField string, cursor interface{}, size uint64, ctx context.Context,
		) ([]Model_T, error)
		FindPrevious(
			cursorField string, cursor interface{}, size uint64, ctx context.Context,
		) ([]Model_T, error)
	}

	IOffsetPaginationRepository[Model_T any] interface {
		FindOffset(
			offset uint64, size uint64, ctx context.Context,
		) ([]Model_T, error)
	}

	IPaginateClonableRepository[Model_T any] interface {
		IPaginationRepository[Model_T]
		Clone() IPaginationRepository[Model_T]
	}

	IPaginationRepository[Model_T any] interface {
		IFilterMethods[Model_T]
		IProjector[Model_T]
		ICursorPaginationRepository[Model_T]
		IOffsetPaginationRepository[Model_T]
		//ICRUDRepository[Model_T]
	}
)
