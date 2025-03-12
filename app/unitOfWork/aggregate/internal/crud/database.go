package internal

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/api/crud"
)

type (
	Database[Read_Entity_T any, Local_Storage_Unit_Entity_T any] struct {

		//storage.IDBStorageQueryExecutor[Local_Storage_Unit_Entity_T]
		executor_generator[Read_Entity_T, Local_Storage_Unit_Entity_T]
		ReadQueryBuilderGenerator         query.IReadQueryBuilderGenerator
		ReadRelationQueryBuilderGenerator relation.IReadRelationQueryBuilderGenerator
	}
)

func (this *Database[Read_Entity_T, Local_Storage_Unit_Entity_T]) Read(
	initFn func(queryBuilder query.IQueryBuilder),
) query.IGenericQueryExecutor[Read_Entity_T] {

	if initFn == nil {

		panic("query builder initialize funtion could not be nil")
	}

	query := this.ReadQueryBuilderGenerator.NewQueryBuilder()

	initFn(query)

	//return read.NewExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T](this.Stu, query)
	return this.NewQueryExecutor(query)
}

// func (this *Database[Read_Entity_T, Local_Storage_Unit_Entity_T]) _read(
// 	initFn func(queryBuilder query.IQueryBuilder), defaultQueryBuilder query.IQueryBuilder,
// ) query.IGenericQueryExecutor[Read_Entity_T] {

// }

func (this *Database[Read_Entity_T, Local_Storage_Unit_Entity_T]) ByRelations(
	realations ...relation.IDBRelationInitiator,
) crud.IAggregateReader[Read_Entity_T] {

	queryBuilder := this.ReadRelationQueryBuilderGenerator.NewRelationQuery()

	queryBuilder.PushRelations(
		realations...,
	)

	// return read.NewRelationExecutor[Read_Entity_T, Local_Storage_Unit_Entity_T](
	// 	this.Stu, queryBuilder,
	// )

	return this.NewRelationReadQueryExecutor(queryBuilder)
}

// func (this *Database[Read_Entity_T, Local_Storage_Unit_Entity_T]) Create(
// 	list []Local_Storage_Unit_Enfunc (this *)tity_T, ctx context.Context,
// ) {

// }

// func (this *Database[Read_Entity_T, Local_Storage_Unit_Entity_T]) Update(
// 	filterFunc func(filter query.IFilterExpression),
// ) {

// }
