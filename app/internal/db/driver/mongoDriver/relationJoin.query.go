package mongoDriver

import (
	mongoRelation "app/internal/db/driver/mongoDriver/relation"
	"app/internal/db/query"
	"app/internal/db/relation"
)

type (
	RelationQueryBuilder[Entity_T any, Source_Repo_Entity any] struct {
		//DelegatorQueryBuilder[Entity_T]
		MongoDBDelegator[Entity_T, Source_Repo_Entity]
	}
)

func (this *RelationQueryBuilder[Entity_T, Source_Repo_Entity]) AggregateRelations(
	relation ...relation.IDBRelationshipNavigator[mongoRelation.Query_Type],
) query.IQueryBuilder[Entity_T] {

	queryBuilder := this.Query()

	return queryBuilder
}
