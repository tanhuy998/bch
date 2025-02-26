package mongoQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/filter"
	"app/internal/db/driver/mongoDriver/sort"
	"app/internal/db/driver/mongoDriver/transform"
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	MongoAggregateQueryBuilder struct {
		mongo_pipeline
		projection map[string]uint
	}
)

func (this *MongoAggregateQueryBuilder) GetPipeline() []interface{} {

	return this.mongo_pipeline.P
}

func (this *MongoAggregateQueryBuilder) initProjection() {

	if this.projection != nil {

		return
	}

	this.projection = make(map[string]uint)
}

func (this *MongoAggregateQueryBuilder) Join(another string, fn func(query.IJoinField)) query.ISubQueryBuilder {

	joinIntializer := &JoinOperationInitializer{
		From: another,
	}

	fn(joinIntializer)

	var ops []interface{}

	if joinIntializer.is_unwind {

		/*
			when the join operation has data limit less than or equal 1
			unwind looked up nested documents for explicit query on joined collection

			example aggregate pipeline:
			[
				{
					$lookup: {
						"form": "anotherCollecction",
						"localField": "id",
						"foreignField": "fID",
						"as": "tests"
					}
				},
				{
					$set: {
						test
					}
				}
			]
		*/

		ops = make([]interface{}, 2)

		ops[1] = bson.D{
			{"$unwind", joinIntializer.Alias},
		}

	} else {

		ops = make([]interface{}, 1)
	}

	ops[0] = bson.D{
		{"$lookup", joinIntializer},
	}

	this.mongo_pipeline.PushStages(ops...)

	return this
}

func (this *MongoAggregateQueryBuilder) Filter(fn query.FilterFunc) query.ISubQueryBuilder {

	filter := filter.NewFilterGenerator()

	fn(filter)

	this.mongo_pipeline.PushStages(
		bson.D{
			{"$match", filter.Get()},
		},
	)

	return this
}

func (this *MongoAggregateQueryBuilder) Select(fields ...string) query.ISubQueryBuilder {

	this.initProjection()

	for _, name := range fields {

		this.projection[name] = 1
	}

	return this
}

func (this *MongoAggregateQueryBuilder) ExcludeFields(fields ...string) query.ISubQueryBuilder {

	this.initProjection()

	for _, name := range fields {

		this.projection[name] = 0
	}

	return this
}

func (this *MongoAggregateQueryBuilder) Done() {

	this.mergeProjection()
}

func (this *MongoAggregateQueryBuilder) mergeProjection() {

	if this.projection == nil {

		return
	}

	this.PushStages(
		bson.D{
			{"$project", this.projection},
		},
	)
}

func (this *MongoAggregateQueryBuilder) Transform(fn query.DataTransformFunc) query.ISubQueryBuilder {

	transformer := transform.NewDataTransformer()

	fn(transformer)

	this.mongo_pipeline.PushStages(
		transformer.GetQuery()...,
	)

	return this
}

func (this *MongoAggregateQueryBuilder) SortOrder(fn query.SortFunc) query.ISubQueryBuilder {

	initializer := sort.NewSortInitializer()

	fn(initializer)

	this.mongo_pipeline.PushStages(
		bson.D{
			{"$sort", initializer},
		},
	)

	return this
}

func (this *MongoAggregateQueryBuilder) Limit(num uint) query.ISubQueryBuilder {

	this.mongo_pipeline.PushStages(
		bson.D{
			{"$limit", num},
		},
	)

	return this
}

func (this *MongoAggregateQueryBuilder) Clone() *MongoAggregateQueryBuilder {

	ret := new(MongoAggregateQueryBuilder)

	ret.P = this.P[:]

	return ret
}
