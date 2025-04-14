package relationQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/driver/mongoDriver/transform"
	"app/internal/db/relation"
)

type (
	RelationLocalNavigator struct {
		/*
			foreign reference
		*/
		ptr_foreign_navigator *RelationForeignNavigator
		/*
			local query builder, implement query.ICLonableQueryBuilder
		*/
		queryBuilder.MongoAggregateQueryBuilder
	}
)

// func NewRelationLocalNaviagator(
// 	localQueryBuilder *relation_dispatcher,
// 	foreignNavigator *RelationForeignNavigator,
// ) *RelationLocalNavigator {

// 	ret := new(RelationLocalNavigator)

// 	ret.ptr_foreign_navigator = foreignNavigator
// 	ret.MongoAggregateQueryBuilder = localQueryBuilder.MongoAggregateQueryBuilder

// 	return ret
// }

func (this *RelationLocalNavigator) Transform(fn relation.RelationDataTransformFunc) relation.IRelationLocalNavigator {

	transformer := transform.NewRelationDataTransformer(
		this.ptr_foreign_navigator,
	)

	fn(transformer)

	(this.MongoAggregateQueryBuilder).PushStages(
		transformer.GetQuery()...,
	)

	return this
}
