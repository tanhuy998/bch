package relationQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/relation"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	relation_dispatcher struct {
		*queryBuilder.MongoAggregateQueryBuilder
	}
)

func NewRelationDispatcher(ref *queryBuilder.MongoAggregateQueryBuilder) *relation_dispatcher {

	if ref == nil {

		panic(
			`relation_dispatcher must point to an "app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder\".MongoAggregateQueryBuilder`,
		)
	}

	ret := new(relation_dispatcher)

	ret.MongoAggregateQueryBuilder = ref

	return ret
}

func (this *relation_dispatcher) PushRelations(
	relation ...relation.IDBRelationInitiator,
) {

	this._dispatch(relation)
}

func (this *relation_dispatcher) _dispatch(
	relations []relation.IDBRelationInitiator,
) {

	for _, initiator := range relations {

		// foreignNavigator := NewRelationForeignNavigator()
		// foreignNavigator.join_op.From = initiator.GetDBStorageUnitName()

		// this.PushStages(
		// 	bson.D{
		// 		{"$lookup", &foreignNavigator.join_op},
		// 	},
		// )

		// initFn(&foreignNavigator.join_op)

		// initiator.ResolveRelation(
		// 	NewRelationLocalNaviagator(
		// 		this, foreignNavigator,
		// 	),
		// 	foreignNavigator,
		// )

		if v, ok := initiator.(relation.IDBRelationBeforeJoinInterceptor); ok {

			v.InterceptBeforeJoin(this)
		}

		relationResolver := NewRelationResolver(initiator)

		foreignInitializer := relationResolver.GetForeignInitializer()

		this.PushStages(
			bson.D{
				{"$lookup", foreignInitializer},
			},
		)

		switch v := initiator.(type) {
		case relation.IDBRelationForeignInitializer:
			v.InitializeForeign(foreignInitializer)
		default:
			initFn := initiator.GetRelationInitFunc()

			if initFn == nil {

				panic("relation init function could not be nil")
			}

			initFn(foreignInitializer)
		}

		relationResolver.Resolve()

		this.MongoAggregateQueryBuilder.PushStages(
			relationResolver.local_navigator.P...,
		)

		if v, ok := initiator.(relation.IDBRelationAfterJoinInterceptor); ok {

			v.InterceptAfterJoin(this)
		}
	}
}
