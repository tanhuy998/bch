package queryBuilder

import (
	"app/internal/db/query"
	"app/internal/db/storage"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	JoinOperationDispatcher struct {
		*MongoAggregateQueryBuilder // reference to the main query builder
		JoinInitializer             *JoinOperationInitializer
	}
)

func NewJoinUnwindableDispatcher(ref *MongoAggregateQueryBuilder) *JoinOperationDispatcher {

	ret := &JoinOperationDispatcher{
		MongoAggregateQueryBuilder: ref,
		JoinInitializer:            NewJoinOperationQueryBuilder(),
	}

	return ret
}

func (this *JoinOperationDispatcher) UnwindForeign() query.IQueryBuilder {

	this.MongoAggregateQueryBuilder.Unwind(
		this.JoinInitializer.Alias,
	)

	return this.MongoAggregateQueryBuilder
}

func (this *JoinOperationDispatcher) DoJoin(
	another storage.IDBStorageIdentifier, fn func(query.IJoinField),
) *JoinOperationDispatcher {

	this.JoinInitializer.From = another.GetDBStorageUnitName()

	fn(this.JoinInitializer)

	this.MongoAggregateQueryBuilder.PushStages(
		"$lookup", this.JoinInitializer,
	)

	return this
}

func (this *JoinOperationDispatcher) AsInnerJoin() {

	this.MongoAggregateQueryBuilder.PushStages(
		bson.D{
			{
				"$unwind", bson.D{
					{"path", fmt.Sprintf(`$%s`, this.JoinInitializer.Alias)},
					{"preserveNullAndEmptyArrays", false},
				},
			},
		},
	)
}

func (this *JoinOperationDispatcher) AsLeftJoin() {

	this.UnwindForeign()
}
