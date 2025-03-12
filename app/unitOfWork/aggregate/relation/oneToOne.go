package aggregateRelation

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	Query_Type = []interface{}
)

type (
	OneToOneWith[Foreign_Entity_T any] struct {
		abstract_relation[Foreign_Entity_T]
	}
)

// func (this *OneToOneWith[Foreign_Entity_T]) ResolveRelationQuery(initializer relation.IDBRelationQueryMetadata) Query_Type {

// 	initializer.SetLimit(1)
// 	alias := initializer.GetAliasName()

// 	ret := []interface{}{
// 		bson.D{
// 			{"$lookup", initializer},
// 		},
// 		bson.D{
// 			{
// 				"$set", bson.D{
// 					{
// 						alias, bson.D{
// 							{
// 								"$arrayElemAt", bson.A{
// 									//"$commandGroup", 0,
// 									fmt.Sprintf("$%s", alias), 0,
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	return ret[:]
// }

func (this *OneToOneWith[Foreign_Entity_T]) ResolveRelation(
	refQueryBuilder query.IQueryBuilder, foreignInitializer relation.IDBRelationNavigator,
) {

	foreignInitializer.SetLimit(1)
	alias := foreignInitializer.GetAliasName()

	refQueryBuilder.Transform(
		func(transform query.IDataTransformer) {

			transform.Set(alias).Value(
				bson.D{
					{
						"$arrayElemAt", bson.A{
							//"$commandGroup", 0,
							fmt.Sprintf("$%s", alias), 0,
						},
					},
				},
			)
		},
	)
}

func (this *OneToOneWith[Foreign_Entity_T]) GetDBRelationKind() string {

	return "one_to_one"
}
