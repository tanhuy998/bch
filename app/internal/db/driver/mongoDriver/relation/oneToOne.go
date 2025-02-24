package mongoRelation

import (
	"app/internal/db/relation"
	"encoding/json"
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
	fmt.Println("one to one with:", initializer.GetAliasName())
	initializer.SetLimit(1)
	alias := initializer.GetAliasName()

	j, _ := json.Marshal(initializer)

	fmt.Println("--------------", string(j))

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
