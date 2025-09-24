package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	opLog "app/unitOfWork/operationLog"
	"context"

	"github.com/google/uuid"
)

type (
	IPaginationFilter repositoryAPI.FilterFunc

	ICursorPaginator[Cursor_T comparable] interface {
		SetCursor(c Cursor_T)
		SetCursorNilValue(v Cursor_T)
		SetCursorDirection(paginateServicePort.CursorDirection)
		CursorFirst()
	}

	IOffsetPaginator interface {
		SetOffset(offset int64)
		SetSize(size uint64)
	}

	IPagaintorProjection = paginateServicePort.IPaginateProjector

	IPaginatorFilter interface {
		ApplyFilter(fn repositoryAPI.FilterFunc)
	}

	IFilterablePaginator interface {
		ApplyPaginateFilter(filterGenerator repositoryAPI.IFilterGenerator)
	}

	IOptionablePaginatorInitializer[Cusor_T comparable] interface {
		IAbstractOptionableOffsetPaginatorInitializer
		ICursorPaginator[Cusor_T]
	}

	IAbstractOptionableOffsetPaginatorInitializer interface {
		IOffsetPaginator
		IPaginatorFilter
		IPagaintorProjection
		getLogRecorder() *opLog.LogRecorder
	}
)

type (
	IExecutablePaginator[Entity_T any] interface {
		repositoryAPI.IOffsetPaginationRepository[Entity_T]
		repositoryAPI.ICursorPaginationRepository[Entity_T]
		repositoryAPI.ISelfStatisticable
		opLog.ILogRecorder
		get_paginator_initializer() IAbstractOptionableOffsetPaginatorInitializer
		GetPageSize() uint64
		GetOffset() int64
		Free()
	}

	IExecutableCursorPaginator[Cursor_T comparable] interface {
		IAbstractCursorPaginator
		GetCursor() *Cursor_T
		IsPrevDirection() bool
	}

	IGeneralExecutableCursorPaginator interface {
		GetCursor() interface{}
		IsPrevDirection() bool
	}

	IAbstractCursorPaginator interface {
		HasCursor() bool
	}
)

type (
	IExecutablePaginatorGenerator[Data_T any, Cursor_T comparable] interface {
		ByDefaultPaginate(
			tenantUUID uuid.UUID, options ...PaginationOption[Cursor_T],
		) IExecutablePaginator[Data_T]
		ByCustomPaginator(
			tenantUUID uuid.UUID, customPaginator paginateServicePort.IPaginator[Cursor_T],
		) IExecutablePaginator[Data_T]
	}
)

type (
	INavigator[Data_T any /*ICursorEntity[Cursor_T] */, Cursor_T comparable] interface {
		GetData() []Data_T
		QueryDocumentCount() int64
		GetFirstCursor() *Cursor_T
		GetLastCursor() *Cursor_T
		CurrentPageNumber() int64
		GetPageSize() uint64
		GetRequestDataSize() uint64
	}
)

type (
	IPaginateInitiator[Entity_T paginateServicePort.ICursorEntity[Cursor_T], Cursor_T comparable] interface {
		Paginate(
			tenantUUID uuid.UUID, options ...PaginationOption[Cursor_T],
		) IPaginateDataRetriever[Entity_T, Cursor_T]
		UseCustomPaginator(
			tenantUUID uuid.UUID, customPaginator paginateServicePort.IPaginator[Cursor_T],
		) IPaginateDataRetriever[Entity_T, Cursor_T]
	}
)

type (
	IPaginateDataRetriever[Entity_T any, Cursor_T comparable] interface {
		Retrieve(ctx context.Context) ([]Entity_T, error)
	}

	IPaginateDataNavigatorRetriever[Entity_T paginateServicePort.ICursorEntity[Cursor_T], Cursor_T comparable] interface {
		ExportAsNavigator(ctx context.Context) (INavigator[Entity_T, Cursor_T], error)
	}
)
