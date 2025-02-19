package query

import (
	paginateServicePort "app/port/paginate"
	"context"
)

type (
	PaginateInitFunc[Cursor_T comparable] func() paginateServicePort.IPaginator[Cursor_T]

	IQueryPaginate[Entity_T any] interface {
		Paginate(context.Context, PaginateInitFunc[interface{}]) ([]Entity_T, error)
	}
)
