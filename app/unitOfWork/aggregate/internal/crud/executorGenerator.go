package internal

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	repositoryAPI "app/repository/api"
	"app/unitOfWork/aggregate/api/crud"
	"app/unitOfWork/aggregate/internal/crud/read"
)

type (
	executor_generator[Read_Entity_T, Local_Storage_Unit_Entity_T any] struct {
		/*
			repository struct inplements this interface
		*/
		Stu repositoryAPI.ICRUDRepository[Local_Storage_Unit_Entity_T]
	}
)

func (this *executor_generator[Read_Entity_T, Local_Storage_Unit_Entity_T]) NewQueryExecutor(
	queryBuilder query.IQueryBuilder,
) query.IGenericQueryExecutor[Read_Entity_T] {

	return read.NewExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T](
		this.Stu, queryBuilder,
	)
}

func (this *executor_generator[Read_Entity_T, Local_Storage_Unit_Entity_T]) NewRelationReadQueryExecutor(
	queryBuilder relation.IReadRelationQueryBuilder,
) crud.IAggregateReader[Read_Entity_T] {

	return read.NewRelationExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T](
		this.Stu, queryBuilder,
	)
}
