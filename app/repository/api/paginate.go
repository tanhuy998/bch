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

	IClonableRepository[Model_T any] interface {
		Clone() IPaginationRepository[Model_T]
	}

	IPaginationRepository[Model_T any] interface {
		IFilterMethods[Model_T]
		IProjectionMethods[Model_T]
		ICursorPaginationRepository[Model_T]
		IOffsetPaginationRepository[Model_T]
		IClonableRepository[Model_T]
	}
)
