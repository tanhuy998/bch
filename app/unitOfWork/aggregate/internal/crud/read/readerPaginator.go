package read

import (
	"app/internal/db/query"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	paginateServicePort "app/port/paginate"
	"context"
	"fmt"
)

type (
	ReaderPaginateExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T any] struct {
		DatabaseReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]
	}
)

func (this *ReaderPaginateExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) Paginate(
	paginator paginateServicePort.IPaginator[interface{}], ctx context.Context,
) ([]Read_Entity_T, error) {

	if paginator == nil {

		return nil, libError.NewInternal(
			fmt.Errorf("could not apply pagination query, nil paginator given"),
		)
	}

	this.resolvePaginateOperators(
		paginator,
	)

	ret := make([]Read_Entity_T, 0)

	err := this.stu.ToSlice(&ret, this.query_builder, ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *ReaderPaginateExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) resolvePaginateOperators(
	paginator paginateServicePort.IPaginator[interface{}],
) {

	cursor_p, isCursor := (paginator).(paginateServicePort.ICursorPaginator[interface{}])
	nil_cursor_p, isNillableCursor := (paginator).(paginateServicePort.ICursorNillablePaginator[interface{}])

	var matchCursorQuery bool = isCursor && (cursor_p.GetCursor() != nil || isNillableCursor && cursor_p.GetCursor() != nil_cursor_p.CursorNilValue())

	if matchCursorQuery {

		this.resovleCursorPaginateOperator(cursor_p)

		return
	}

	pageNumber := paginator.GetPageNumber()
	pageSize := paginator.GetPageSize()

	this.query_builder.SortOrder(
		func(sorter query.ISortInitializer) {

			sorter.Field("_id").Descending()
		},
	)
	this.query_builder.Skip(pageNumber)
	this.query_builder.Limit(
		libCommon.Ternary[uint64](pageSize == 0, 1, pageSize),
	)
}

func (this *ReaderPaginateExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) resovleCursorPaginateOperator(
	paginator paginateServicePort.ICursorPaginator[interface{}],
) {

	isCursorPrevDir := paginator.GetCursorDirection() == paginateServicePort.CURSOR_DIRECTION_PREVIOUS

	if isCursorPrevDir {

		this.resolvePrevOperators(paginator)
	}

	this.resolveNextOperators(paginator)
}

func (this *ReaderPaginateExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) resolveNextOperators(
	paginator paginateServicePort.ICursorPaginator[interface{}],
) {

	this.query_builder.SortOrder(
		func(sorter query.ISortInitializer) {

			sorter.Field("_id").Descending()
		},
	)

	this.query_builder.Filter(
		func(filter query.IFilterExpression) {

			filter.Field("_id").LessThan(paginator.GetCursor())
		},
	)

	pageSize := paginator.GetPageSize()

	this.query_builder.Limit(
		libCommon.Ternary(pageSize == 0, 1, pageSize),
	)
}

func (this *ReaderPaginateExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) resolvePrevOperators(
	paginator paginateServicePort.ICursorPaginator[interface{}],
) {

	this.query_builder.SortOrder(
		func(sorter query.ISortInitializer) {

			sorter.Field("_id").Ascending()
		},
	)

	this.query_builder.Filter(
		func(filter query.IFilterExpression) {

			filter.Field("_id").GreaterThan(paginator.GetCursor())
		},
	)

	pageSize := paginator.GetPageSize()

	this.query_builder.Limit(
		libCommon.Ternary(pageSize == 0, 1, pageSize),
	)

	this.query_builder.SortOrder(
		func(sorter query.ISortInitializer) {

			sorter.Field("_id").Descending()
		},
	)
}
