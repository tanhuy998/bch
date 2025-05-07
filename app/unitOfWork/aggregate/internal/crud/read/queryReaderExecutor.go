package read

import (
	"app/internal/db/query"
	"app/unitOfWork/aggregate/api/crud"
)

type (
	QueryReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T any] struct {
		ReaderPagianteExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]
	}
)

func (this *QueryReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) Read(
	initFn crud.QueryBuilderFunc,
) (reader query.IGenericQueryExecutor[Read_Entity_T]) {

	reader = this

	if initFn == nil {

		//panic("aggregate read init function must not be nil")

		return
	}

	initFn(this.query_builder)

	return
}
