package mongoRelation

import (
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

func (this *OneToOneWith[Foreign_Entity_T]) ResolveRelationQuery(initializer relation.IDBRelationQueryMetadata) Query_Type {

	initializer.SetLimit(1)
	alias := initializer.GetAliasName()

	ret := []interface{}{
		bson.D{
			{"$lookup", initializer},
		},
		bson.D{
			{
				"$set", bson.D{
					{
						alias, bson.D{
							{
								"$arrayElemAt", bson.A{
									//"$commandGroup", 0,
									fmt.Sprintf("$%s", alias), 0,
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
