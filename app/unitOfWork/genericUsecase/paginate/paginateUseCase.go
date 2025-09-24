package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	"context"
	"fmt"

	"github.com/google/uuid"
)

const (
	DEFAULT_CURSOR_FIELD = "_id"
)

type (
	PaginationOption[Cursor_T comparable] func(paginator IOptionablePaginatorInitializer[Cursor_T])
)

type (
	/*
		PaginateUseCase unit of work setups pagination handling use case for list data retrieval.
		The Pagination use case uses both cursor and offset pagination method.
		The priority of applying the exact method is: cursor pagination (if cursor value is provided and not equal cursor's nil value)
		-> offset (when cursor pagination prerequisites are not met).
	*/
	PaginateUseCase[
		// require a repository of any types that implements repsitoryAPI.IPaginateClonableRepository
		// in order to provide a query builder for the pagintion process.
		Repository_T repositoryAPI.IPaginateClonableRepository[Entity_T],
		// the Entity type correspoding the storage unit that the repository is mapped to.
		Entity_T any,
		// the type of cursor for the PaginateUseCase to build the query and must be
		// a comparable type. Cursor type ust be provide to ensure data consitency among services.
		Cursor_T comparable,
	] struct {
		//logger
		//internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]
		//paginate_executor_t[Repository_T, Entity_T, Cursor_T]
		//Repository Repository_T // repositoryAPI.IPaginateClonableRepository[Entity_T]
		paginate_dispatcher_t[Repository_T, Entity_T, Cursor_T]
	}
)

func (this *PaginateUseCase[Repository_T, Entity_T, Cursor_T]) Paginate(
	tenantUUID uuid.UUID, ctx context.Context, options ...PaginationOption[Cursor_T],
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	_p := this.generateExecutablePaginatorByDefault(
		tenantUUID, options,
	)

	return this._dispatch(_p.paginator_t, ctx)
}

func (this *PaginateUseCase[Repository_T, Entity_T, Cursor_T]) UseCustomPaginator(
	tenantUUID uuid.UUID, customPaginator paginateServicePort.IPaginator[Cursor_T], ctx context.Context,
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	_p := this.generateExecutablePaginatorByCustomPaginator(tenantUUID, customPaginator)

	return this.dispatch(_p, ctx)
}
