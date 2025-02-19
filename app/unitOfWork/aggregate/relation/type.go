package aggregateRelation

import (
	mongoRelation "app/internal/db/driver/mongoDriver/relation"
	"app/internal/db/relation"
)

type (
	AbstractRelation = relation.IDBRelationshipNavigator[mongoRelation.Query_Type]

	OneToOneWith[Repo_Entity_T any] struct {
		mongoRelation.OneToOneWith[Repo_Entity_T]
	}
)
