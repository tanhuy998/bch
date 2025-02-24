package mongoDriver

import (
	mongoRelation "app/internal/db/driver/mongoDriver/relation"
	"app/internal/db/query"
	"app/internal/db/relation"
	libCommon "app/internal/lib/common"
)

type (
	RelationDelegator[Entity_T any, Source_Repo_Entity any] struct {
		MongoDBDelegator[Entity_T, Source_Repo_Entity]
		//RelationDelegatorQueryBuilder[Entity_T]
	}
)

func (this *RelationDelegator[Entity_T, Source_Repo_Entity_T]) Clone() *RelationDelegator[Entity_T, Source_Repo_Entity_T] {

	return libCommon.PointerPrimitive(*this)
}

func (this *RelationDelegator[Entity_T, Source_Repo_Entity_T]) Query() query.IQueryBuilder[Entity_T] {

	return NewRelationDelegatorQueryBuilder[Entity_T](this.Stu.GetStorageUnit())
}

func (this *RelationDelegator[Entity_T, Source_Repo_Entity_T]) Exec(query query.IQueryBuilder[Entity_T]) query.IQueryExecutor[Entity_T] {

	executor := query.NewExecutor(this)

	return executor
}

func (this *RelationDelegator[Entity_T, Source_Repo_Entity]) AggregateRelations(
	relations ...relation.IDBRelationshipNavigator[mongoRelation.Query_Type],
) query.IQueryBuilder[Entity_T] {

	return NewRelationDelegatorQueryBuilder[Entity_T](this.Stu.GetStorageUnit()).AggregateRelations(relations...)
}
