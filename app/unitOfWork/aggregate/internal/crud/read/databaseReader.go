package read

import (
	"app/internal/db/query"
	"app/internal/db/storage"
	"context"
)

type (
	DatabaseReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T any] struct {
		stu           storage.IDBStorageQueryExecutor[Local_Storage_Unit_Entity_T]
		query_builder query.IQueryBuilder
	}
)

func (this *DatabaseReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) First(
	ctx context.Context,
) (*Read_Entity_T, error) {

	ret := new(Read_Entity_T)

	err := this.stu.First(ret, this.query_builder, ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *DatabaseReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T]) All(
	ctx context.Context,
) ([]Read_Entity_T, error) {

	ret := make([]Read_Entity_T, 0)

	err := this.stu.ToSlice(ret, this.query_builder, ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}
