package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	PaginationOption[Cursor_T comparable] func(paginator IPaginatorInitializer[Cursor_T])
)

type (
	/*
		PaginateUseCase unit of work setups pagination handling use case for list data retrieval.
		The Pagination use case uses both cursor and offset pagination method.
		The priority of applying the exact method is: cursor pagination (if cursor value is provided and not equal cursor's nil value)
		-> offset (when cursor pagination prerequisites are not met).
	*/
	PaginateUseCase[
		// require a repository that implements repsitoryAPI.IPaginateClonableRepository of a particular type
		// in order to provide a query builder for the pagintion process.
		Repository_T repositoryAPI.IPaginateClonableRepository[Entity_T],
		// the type of data that will be binded to the query result
		Entity_T any,
		// the type of cursor for the PaginateUseCase to build the query and must be
		// a comparable type. Cursor type ust be provide to ensure data consitency among services.
		Cursor_T comparable,
	] struct {
		logger
		PaginateRepo Repository_T // repositoryAPI.IPaginateClonableRepository[Entity_T]
	}
)

func (this *PaginateUseCase[Repository_T, Entity_T, Cursor_T]) Paginate(
	tenantUUID uuid.UUID, ctx context.Context, options ...PaginationOption[Cursor_T],
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	paginator := NewPaginator[Entity_T, Cursor_T](this.PaginateRepo)

	for _, fn := range options {

		fn(paginator)
	}

	paginator.IPaginationRepository.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	)

	return this.exec(paginator, ctx)
}

func (this *PaginateUseCase[Repository_T, Entity_T, Cursor_T]) exec(
	paginator *paginator[Entity_T, Cursor_T], ctx context.Context,
) (ret []Entity_T, err error) {

	defer paginator.Free()

	if !paginator.HasCursor() {

		defer func() {
			this.logOffset(paginator.Offset, paginator.Size, err, ctx)
		}()

		return paginator.FindOffset(
			paginator.Offset, paginator.Size, ctx,
		)
	}

	if paginator.IsPrev {

		defer func() {
			this.logCursor(paginator.Cursor, paginator.Size, paginateServicePort.CURSOR_DIRECTION_PREVIOUS, err, ctx)
		}()

		return paginator.FindPrevious(
			"_id", paginator.Cursor, paginator.Size, ctx,
		)
	}

	defer func() {
		this.logCursor(paginator.Cursor, paginator.Size, paginateServicePort.CURSOR_DIRECTION_NEXT, err, ctx)
	}()

	return paginator.FindNext(
		"_id", paginator.Cursor, paginator.Size, ctx,
	)
}

func (this *PaginateUseCase[Repository_T, Entity_T, Cursor_T]) UseCustomPaginator(
	tenantUUID uuid.UUID, customePaginator paginateServicePort.IPaginator[Cursor_T], ctx context.Context,
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	_p := NewPaginator[Entity_T, Cursor_T](this.PaginateRepo)

	_p.IPaginationRepository.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			if v, ok := (customePaginator).(paginateServicePort.IFilterablePaginator); ok {

				v.ApplyPaginateFilter(filter)
			}

			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	)

	if v, ok := (customePaginator).(paginateServicePort.IProjectionPaginator); ok {

		v.ApplyPaginateProjection(_p)
	}

	if v, ok := (customePaginator).(paginateServicePort.ICursorNillablePaginator[Cursor_T]); ok {

		if v.CursorNilValue() != nil {

			_p.SetCursorNilValue(*v.CursorNilValue())
		}
	}

	if v, ok := (customePaginator).(paginateServicePort.ICursorPaginator[Cursor_T]); ok {

		if v.GetCursor() != nil {

			_p.SetCursor(*v.GetCursor())
			_p.SetCursorDirection(v.GetCursorDirection())
			_p.CursorFirst()
		}
	}

	_p.SetOffset(customePaginator.GetPageNumber())
	_p.SetSize(customePaginator.GetPageSize())

	return this.exec(_p, ctx)
}
