package mongoQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/transform"
	"app/internal/db/relation"
)

type (
	JoinOperationPipelineQueryBuilder struct {
		relation_navigator relation.IDBRelationNavigator
		*MongoAggregateQueryBuilder
	}
)

func NewRelationQueryBuilder(
	ref *MongoAggregateQueryBuilder,
	relationNavigator relation.IDBRelationNavigator,
) *JoinOperationPipelineQueryBuilder {

	ret := new(JoinOperationPipelineQueryBuilder)

	ret.relation_navigator = relationNavigator
	ret.MongoAggregateQueryBuilder = ref

	return ret
}

func (this *JoinOperationPipelineQueryBuilder) Transform(fn relation.RelationDataTransformFunc) {

	transformer := transform.NewRelationDataTransformer(
		this.relation_navigator,
	)

	fn(transformer)

	this.MongoAggregateQueryBuilder.PushStages(
		transformer.GetQuery()...,
	)
}
