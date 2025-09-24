package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	"context"
)

type (
	paginate_dispatcher_t[
		// require a repository of any types that implements repsitoryAPI.IPaginateClonableRepository
		// in order to provide a query builder for the pagintion process.
		Repository_T repositoryAPI.IPaginateClonableRepository[Entity_T],
		// the Entity type correspoding the storage unit that the repository is mapped to.
		Entity_T any,
		// the type of cursor for the PaginateUseCase to build the query and must be
		// a comparable type. Cursor type ust be provide to ensure data consitency among services.
		Cursor_T comparable,
	] struct {
		//logger_t
		//Repository Repository_T
		internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]
	}
)

func (this *paginate_dispatcher_t[Repository_T, Entity_T, Cursor_T]) _dispatch(
	paginator paginator_t[Entity_T, Cursor_T], ctx context.Context,
) ([]Entity_T, error) {

	//defer paginator.Free()
	//defer this.writeDebugLogs(paginator, ctx)

	return this.dispatch(
		&executable_paginator[Entity_T, Cursor_T]{
			paginator_t: paginator,
		},
		ctx,
	)
}

func (this *paginate_dispatcher_t[Repository_T, Entity_T, Cursor_T]) dispatch(
	paginator IExecutablePaginator[Entity_T], ctx context.Context,
) (ret []Entity_T, err error) {

	defer paginator.Free()
	defer this.writeDispatchDebugLogs(paginator, ctx)

	switch _p := any(paginator).(type) {
	case IExecutableCursorPaginator[Cursor_T]:
		if !_p.HasCursor() {

			return this.dispatchAsOffsetPaginator(paginator, ctx)
		}

		if _p.IsPrevDirection() {

			defer this.logCursor(
				_p.GetCursor(), paginator.GetPageSize(), paginateServicePort.CURSOR_DIRECTION_PREVIOUS, err, paginator.GetLogRecorder(),
			)

			return paginator.FindPrevious(
				DEFAULT_CURSOR_FIELD, _p.GetCursor(), paginator.GetPageSize(), ctx,
			)
		}

		defer this.logCursor(
			_p.GetCursor(), paginator.GetPageSize(), paginateServicePort.CURSOR_DIRECTION_NEXT, err, paginator.GetLogRecorder(),
		)

		return paginator.FindNext(
			DEFAULT_CURSOR_FIELD, _p.GetCursor(), paginator.GetPageSize(), ctx,
		)
	case IGeneralExecutableCursorPaginator:
		if _p.GetCursor() == nil {

			return this.dispatchAsOffsetPaginator(paginator, ctx)
		}

		if _p.IsPrevDirection() {

			defer this.logCursor(
				_p.GetCursor(), paginator.GetPageSize(), paginateServicePort.CURSOR_DIRECTION_PREVIOUS, err, paginator.GetLogRecorder(),
			)

			return paginator.FindPrevious(
				DEFAULT_CURSOR_FIELD, _p.GetCursor(), paginator.GetPageSize(), ctx,
			)
		}

		defer this.logCursor(
			_p.GetCursor(), paginator.GetPageSize(), paginateServicePort.CURSOR_DIRECTION_NEXT, err, paginator.GetLogRecorder(),
		)

		return paginator.FindNext(
			DEFAULT_CURSOR_FIELD, _p.GetCursor(), paginator.GetPageSize(), ctx,
		)
	default:
		return this.dispatchAsOffsetPaginator(paginator, ctx)
	}

	// if !paginator.HasCursor() {

	// 	defer this.logOffset(
	// 		paginator.GetOffset(), paginator.GetPageSize(), err, paginator.GetLogRecorder(),
	// 	)

	// 	return paginator.FindOffset(
	// 		uint64(paginator.GetOffset()), paginator.GetPageSize(), ctx,
	// 	)
	// }

	// if paginator.IsPrevDirection() {

	// 	defer this.logCursor(
	// 		paginator.GetCursor(), paginator.GetPageSize(), paginateServicePort.CURSOR_DIRECTION_PREVIOUS, err, paginator.GetLogRecorder(),
	// 	)

	// 	return paginator.FindPrevious(
	// 		DEFAULT_CURSOR_FIELD, paginator.GetCursor(), paginator.GetPageSize(), ctx,
	// 	)
	// }

	// defer this.logCursor(
	// 	paginator.GetCursor(), paginator.GetPageSize(), paginateServicePort.CURSOR_DIRECTION_NEXT, err, paginator.GetLogRecorder(),
	// )

	// return paginator.FindNext(
	// 	DEFAULT_CURSOR_FIELD, paginator.GetCursor(), paginator.GetPageSize(), ctx,
	// )
}

func (this *paginate_dispatcher_t[Repository_T, Entity_T, Cursor_T]) dispatchAsOffsetPaginator(
	paginator IExecutablePaginator[Entity_T], ctx context.Context,
) (ret []Entity_T, err error) {

	defer this.logOffset(
		paginator.GetOffset(), paginator.GetPageSize(), err, paginator.GetLogRecorder(),
	)

	return paginator.FindOffset(
		uint64(paginator.GetOffset()), paginator.GetPageSize(), ctx,
	)
}

func (this *paginate_dispatcher_t[Repository_T, Entity_T, Cursor_T]) writeDispatchDebugLogs(
	paginator IExecutablePaginator[Entity_T], execCtx context.Context,
) {

	debugCtx := paginator.GetLogRecorder()

	if debugCtx == nil {

		return
	}

	debugCtx.SetBaseContext(execCtx)
}
