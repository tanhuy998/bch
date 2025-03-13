package aggregateRelation

import (
	"app/internal/db/relation"
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
	queryBuilder relation.IRelationQueryBuilder, foreignInitializer relation.IDBRelationNavigator,
) {

	foreignInitializer.SetLimit(1)
	alias := foreignInitializer.GetAliasName()

	/*
		Decouple the the domain from concrete aspect of mongodb
	*/

	// queryBuilder.Transform(
	// 	func(transform query.IDataTransformer) {

	// 		transform.Set(alias).Value(
	// 			bson.D{
	// 				{
	// 					"$arrayElemAt", bson.A{
	// 						//"$commandGroup", 0,
	// 						fmt.Sprintf("$%s", alias), 0,
	// 					},
	// 				},
	// 			},
	// 		)
	// 	},
	// )

	queryBuilder.Transform(
		func(transform relation.IRelationDataTransformer) {

			transform.Set(alias).AsForeign().FirstElement()
		},
	)
}

/*
for debug log
*/
func (this *OneToOneWith[Foreign_Entity_T]) GetDBRelationKind() string {

	return "one_to_one"
}
