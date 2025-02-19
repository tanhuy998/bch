package mongoRelation

// import (
// 	"app/internal/db/relation"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// type (
// 	ManyToOneWith[Foreign_Entity_T any] struct {
// 		asbtract_relation[Foreign_Entity_T]
// 	}
// )

// func (this *ManyToOneWith[Foreign_Entity_T]) ResolveQuery(initializer relation.IDBRelationQueryMetadata) Query_Type {

// 	initializer.Limit(1)

// 	ret := [1]interface{}{
// 		bson.D{
// 			{"$lookup", initializer},
// 		},
// 	}

// 	return ret[:]
// }
