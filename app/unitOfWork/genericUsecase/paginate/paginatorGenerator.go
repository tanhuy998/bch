package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"

	"github.com/google/uuid"
)

type (
	excutable_paginator_generator_t[
		Repository_T repositoryAPI.IPaginateClonableRepository[Entity_T],
		Entity_T any,
		Cursor_T comparable,
	] internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]
)

func (this *excutable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]) ByDefaultPaginate(
	tenantUUID uuid.UUID, options ...PaginationOption[Cursor_T],
) IExecutablePaginator[Entity_T] {

	return internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T](*this).
		generateExecutablePaginatorByDefault(tenantUUID, options)
}

func (this *excutable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]) ByCustomPaginator(
	tenantUUID uuid.UUID, customPaginator paginateServicePort.IPaginator[Cursor_T],
) IExecutablePaginator[Entity_T] {

	return internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T](*this).
		generateExecutablePaginatorByCustomPaginator(tenantUUID, customPaginator)
}
