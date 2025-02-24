package mongoRelation

import (
	"app/internal/db/relation"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	OneToManyWith[Foreign_Entity_T any] struct {
		abstract_relation[Foreign_Entity_T]
	}
)

func (this *OneToManyWith[Foreign_Entity_T]) ResolveQuery(initializer relation.IDBRelationQueryMetadata) Query_Type {

	//initializer.SetLimit(1)

	ret := [1]interface{}{
		bson.D{
			{"$lookup", initializer},
		},
	}

	return ret[:]
}
