package aggregate

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/api/crud"
	internal "app/unitOfWork/aggregate/internal/crud"
)

type (
	DefaultAggregateRelation[Read_Entity_T, Local_Storage_Unit_Entity_T any] struct {
		internal.Database[Read_Entity_T, Local_Storage_Unit_Entity_T]
		default_relation_query_builder relation.IClonableReadRelationQueryBuilder
	}
)

func (this *DefaultAggregateRelation[Read_Entity_T, Local_Storage_Unit_Entity_T]) ByDefaultRelations(
	relations ...relation.IDBRelationInitiator,
) crud.IAggregateReader[Read_Entity_T] {

	defaultQueryBuilder := this.resolveDefaultRelationsQueryBuilder(relations...)

	return this.Database.NewRelationReadQueryExecutor(defaultQueryBuilder)
}

func (this *DefaultAggregateRelation[Read_Entity_T, Local_Storage_Unit_Entity_T]) resolveDefaultRelationsQueryBuilder(
	relations ...relation.IDBRelationInitiator,
) relation.IClonableReadRelationQueryBuilder {

	if this.default_relation_query_builder != nil {

		return this.default_relation_query_builder.Clone()
	}

	newDefaultqueryBuilder := this.Database.ReadRelationQueryBuilderGenerator.NewRelationQuery()

	newDefaultqueryBuilder.PushRelations(relations...)

	this.default_relation_query_builder = newDefaultqueryBuilder

	return this.default_relation_query_builder.Clone()
}
