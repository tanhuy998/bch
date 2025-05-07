package read

import (
	paginateServicePort "app/port/paginate"
	"context"
)

type (
	ReaderQueryContextExecutor[Read_Entity_T any, Local_Entity_T any] struct {
		ReaderPaginateExecutor[Read_Entity_T, Local_Entity_T]
	}
)

func (this *ReaderQueryContextExecutor[Read_Entity_T, Local_Entity_T]) ToSLiceByContext(
	context context.Context,
) ([]Read_Entity_T, error) {

	switch paginator, ok := context.(paginateServicePort.IPaginator[interface{}]); {
	case ok:
		return this.Paginate(paginator, context)
	default:
		return this.All(context)
	}
}
