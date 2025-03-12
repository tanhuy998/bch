package aggregateRelation

import (
	"app/internal/db/query"
	"app/internal/db/relation"
)

type (
	OneToManyWith[Foreign_Entity_T any] struct {
		abstract_relation[Foreign_Entity_T]
	}
)

// func (this *OneToManyWith[Foreign_Entity_T]) ResolveRelationQuery(initializer relation.IDBRelationQueryMetadata) Query_Type {

// 	//initializer.SetLimit(1)

// 	ret := [1]interface{}{
// 		bson.D{
// 			{"$lookup", initializer},
// 		},
// 	}

//		return ret[:]
//	}
func (this *OneToManyWith[Foreign_Entity_T]) ResolveRelation(
	refQueryBuilder query.IQueryBuilder, foreignInitializer relation.IDBRelationInitiator,
) {

}

func (this *OneToManyWith[Foreign_Entity_T]) GetDBRelationKind() string {

	return "one_to_many"
}
