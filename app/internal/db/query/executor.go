package query

import (
	paginateServicePort "app/port/paginate"
	"context"
)

type (
	IGenericQueryExecutor[Entity_T any] interface {
		First(ctx context.Context) (*Entity_T, error)
		All(ctx context.Context) ([]Entity_T, error)
		Paginate(
			paginator paginateServicePort.IPaginator[interface{}], ctx context.Context,
		) ([]Entity_T, error)
	}
)
