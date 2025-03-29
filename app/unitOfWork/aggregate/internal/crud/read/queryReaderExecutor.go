package read

import "app/internal/db/query"

type (
	QueryReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T any] struct {
		ReaderPagianteExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]
	}
)

func (this *QueryReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) Read(
	initFn func(query query.IQueryBuilder),
) query.IGenericQueryExecutor[Read_Entity_T] {

	if initFn == nil {

		panic("aggregate read init function must not be nil")
	}

	initFn(this.query_builder)

	return this
}
