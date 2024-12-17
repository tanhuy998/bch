package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	opLog "app/unitOfWork/operationLog"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type (
	PaginationOption func(paginator IPaginatorInitializer)
)

type (
	PaginateUseCase[Entity_T any, Cursor_T comparable] struct {
		opLog.OperationLogger
		PaginateRepo repositoryAPI.IPaginateClonableRepository[Entity_T]
	}
)

func (this *PaginateUseCase[Entity_T, Cursor_T]) Paginate(
	tenantUUID uuid.UUID, ctx context.Context, options ...PaginationOption,
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	paginator := NewPaginator(this.PaginateRepo)

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

func (this *PaginateUseCase[Entity_T, Cursor_T]) exec(
	paginator *paginator[Entity_T], ctx context.Context,
) ([]Entity_T, error) {

	if paginator.Cursor == nil {

		this.OperationLogger.PushTrace("paginate", "offset", ctx)

		return paginator.FindOffset(
			paginator.Offset, paginator.Size, ctx,
		)
	}

	if paginator.IsPrev {

		this.OperationLogger.PushTrace("paginate", "cursor_forward", ctx)

		return paginator.FindPrevious(
			"_id", paginator.Cursor, paginator.Size, ctx,
		)
	}

	this.OperationLogger.PushTrace("paginate", "cursor_prev", ctx)

	return paginator.FindNext(
		"_id", paginator.Cursor, paginator.Size, ctx,
	)
}

func (this *PaginateUseCase[Entity_T, Cursor_T]) UseCustomPaginator(
	tenantUUID uuid.UUID, paginator paginateServicePort.IPaginator[Cursor_T], ctx context.Context,
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	_p := NewPaginator(this.PaginateRepo)

	_p.IPaginationRepository.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			if v, ok := (paginator).(paginateServicePort.IFilterablePaginator); ok {

				v.ApplyPaginateFilter(filter)
			}

			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	)

	if v, ok := (paginator).(paginateServicePort.IProjectionPaginator); ok {

		v.ApplyPaginateProjection(_p)
	}

	if paginator.GetCursor() != paginator.CursorNilValue() {

		_p.SetCursor(paginator.GetCursor())
		_p.SetCursorDirection(paginator.GetCursorDirection())
		_p.CursorFirst()
	}

	_p.SetOffset(paginator.GetPageNumber())
	_p.SetSize(paginator.GetPageSize())

	return this.exec(_p, ctx)
}
