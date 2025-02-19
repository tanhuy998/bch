package mongoRelation

import (
	"app/internal/db/relation"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	Query_Type = []interface{}
)

type (
	OneToOneWith[Foreign_Entity_T any] struct {
		asbtract_relation[Foreign_Entity_T]
	}
)

func (this *OneToOneWith[Foreign_Entity_T]) ResolveQuery(initializer relation.IDBRelationQueryMetadata) Query_Type {

	initializer.SetLimit(1)

	ret := [2]interface{}{
		bson.D{
			{"$lookup", initializer},
		},
		bson.D{
			{
				"$set", bson.D{
					{
						initializer.GetAliasName(), bson.D{
							{
								"$arrayElemAt", bson.A{
									"$commandGroup", 0,
								},
							},
						},
					},
				},
			},
		},
	}

	return ret[:]
}
