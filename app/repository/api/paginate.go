package repositoryAPI

import (
	"context"
)

type (
	IPaginateQueryConditionUnitGetter interface {
		QueryConditionUnit() IPaginateQueryConditionUnit
	}

	IPaginateQueryConditionUnit interface {
		AsFilter(fn FilterFunc)
		AsConditionExpression(fn MatchFunc)
	}
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
		IPaginationUnit[Model_T]
		Clone() IPaginationUnit[Model_T]
	}

	IPaginationUnit[Model_T any] interface {
		IStatisticRepsitory
		ISelfStatisticable
		IFilterMethods[Model_T]
		IProjector[Model_T]
		ICursorPaginationRepository[Model_T]
		IOffsetPaginationRepository[Model_T]
		IPaginateQueryConditionUnitGetter
		//IPaginateFilterableUnit
		//ICRUDRepository[Model_T]
	}
)
