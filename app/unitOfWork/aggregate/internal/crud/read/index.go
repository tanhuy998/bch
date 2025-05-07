package read

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/internal/db/storage"
	"app/unitOfWork/aggregate/api/crud"
)

func NewExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T any](
	stu storage.IDBStorageQueryExecutor[Local_Storage_Unit_Entity_T],
	queryBuilder query.IQueryBuilder,
) query.IGenericQueryExecutor[Read_Entity_T] {

	ret := new(QueryReader[Read_Entity_T, Local_Storage_Unit_Entity_T])

	ret.stu = stu
	ret.query_builder = queryBuilder

	return ret
}

func NewRelationExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T any](
	stu storage.IDBStorageQueryExecutor[Local_Storage_Unit_Entity_T],
	queryBuilder relation.IReadRelationQueryBuilder,
) crud.IAggregateReader[Read_Entity_T] {

	ret := new(RelationReaderExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T])

	ret.stu = stu
	ret.query_builder = queryBuilder

	return ret
}
