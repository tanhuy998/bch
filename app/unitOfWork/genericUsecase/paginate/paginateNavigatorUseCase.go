package paginateUseCase

import (
	repositoryAPI "app/repository/api"
	"context"
	"strconv"
)

type (
	PaginateInitFunc[Data_T any, Cursor_T comparable] func(generator IExecutablePaginatorGenerator[Data_T, Cursor_T]) IExecutablePaginator[Data_T]
)

type (
	PaginateNavigatorUseCase[
		Repository_T repositoryAPI.IPaginateClonableRepository[Entity_T],
		Entity_T any, //ICursorEntity[Cursor_T],
		Cursor_T comparable,
	] struct {
		PaginateUseCase[
			Repository_T, Entity_T, Cursor_T,
		]
	}
)

func (this *PaginateNavigatorUseCase[Repository_T, Entity_T, Cursor_T]) ExportAsNavigator(
	ctx context.Context, fn PaginateInitFunc[Entity_T, Cursor_T],
) (n INavigator[Entity_T, Cursor_T], e error) {

	if fn == nil {

		panic("paginate initialization function passed to ExportAsNavigator() must not be nil")
	}

	debug := this.Debug("paginator_data")

	generator := excutable_paginator_generator_t[Repository_T, Entity_T, Cursor_T](
		this.internal_executable_paginator_generator_t,
	)

	exceutablePaginator := fn(&generator)

	data, err := this.dispatch(exceutablePaginator, ctx)

	if err != nil {

		return nil, err
	}

	queryCount, err := exceutablePaginator.SelfStatistic().Count(ctx)

	if err != nil {

		return nil, err
	}

	ret := &paginate_navigator_t[Entity_T, Cursor_T]{
		data: data,
		//is_cursor_pagination: exceutablePaginator.HasCursor(),
		query_total_count: queryCount,
		request_data_size: int64(exceutablePaginator.GetPageSize()),
	}

	switch v, ok := any(exceutablePaginator).(IAbstractCursorPaginator); {
	case ok:
		ret.is_cursor_pagination = v.HasCursor()
		fallthrough
	case v.HasCursor():
		ret.current_page_number = exceutablePaginator.GetOffset()
	}

	// if !exceutablePaginator.HasCursor() {

	// 	ret.current_page_number = exceutablePaginator.GetOffset()
	// }

	debug.Push("retrieved_data_length", strconv.Itoa(len(data)), ctx)

	return ret, nil
}
