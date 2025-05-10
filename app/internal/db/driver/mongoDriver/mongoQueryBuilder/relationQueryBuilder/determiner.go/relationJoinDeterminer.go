package determiner

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/relation"
)

type (
	RelationJoinDeterminer struct {
		dispatcher queryBuilder.JoinOperationDispatcher
	}
)

// func NewRelationJoinOperationDeteminer(
// 	dispatcher *queryBuilder.JoinOperationDispatcher,
// ) *relation_join_determiner {

// 	ret := &relation_join_determiner{
// 		dispatcher: dispatcher,
// 	}
// 	return ret
// }

func (this *RelationJoinDeterminer) SetJoinInitializer(initializer *queryBuilder.JoinOperationInitializer) {

	this.dispatcher.JoinInitializer = initializer
}

func (this *RelationJoinDeterminer) SetLocalQueryBuilder(queryBuilder *queryBuilder.MongoAggregateQueryBuilder) {

	this.dispatcher.MongoAggregateQueryBuilder = queryBuilder
}

func (this *RelationJoinDeterminer) AsInnerJoin() relation.IJoinOperator {

	return inner_join{
		dispatcher: &this.dispatcher,
	}
}

func (this *RelationJoinDeterminer) AsLeftJoin() relation.IJoinOperator {

	return left_join{
		dispatcher: &this.dispatcher,
	}
}

func (this *RelationJoinDeterminer) DetermineJoinOperation(
	initiator relation.IDBRelationInitiator,
) {

	switch d := initiator.(type) {
	case relation.IDBRelationInitiatorJoinOperationDeterminer:

		d.DetermineJoinOperation(this).ApplyJoinOperation()
	}
}
