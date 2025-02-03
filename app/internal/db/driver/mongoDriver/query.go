package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/filter"
	"app/internal/db/driver/mongoDriver/sort"
	"app/internal/db/driver/mongoDriver/transform"
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	mongo_query struct {
		mongo_pipeline
		projection map[string]uint
	}
)

func (this *mongo_query) initProjection() {

	if this.projection != nil {

		return
	}

	this.projection = make(map[string]uint)
}

func (this *mongo_query) Join(another string, fn func(query.IJoinField)) query.ISubQueryBuilder {

	joinIntializer := &join_op{
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

	this.mongo_pipeline.Add(ops...)

	return this
}

func (this *mongo_query) Filter(fn query.FilterFunc) query.ISubQueryBuilder {

	filter := filter.NewFilterGenerator()

	fn(filter)

	this.mongo_pipeline.Add(
		filter.Get(),
	)

	return this
}

func (this *mongo_query) Select(fields ...string) query.ISubQueryBuilder {

	this.initProjection()

	for _, name := range fields {

		this.projection[name] = 1
	}

	return this
}

func (this *mongo_query) ExcludeFields(fields ...string) query.ISubQueryBuilder {

	this.initProjection()

	for _, name := range fields {

		this.projection[name] = 0
	}

	return this
}

func (this *mongo_query) Done() {

	this.mergeProjection()
}

func (this *mongo_query) mergeProjection() {

	if this.projection == nil {

		return
	}

	this.Add(
		bson.D{
			{"$project", this.projection},
		},
	)
}

func (this *mongo_query) Transform(fn query.DataTransformFunc) query.ISubQueryBuilder {

	transformer := transform.NewDataTransformer()

	fn(transformer)

	this.mongo_pipeline.Add(
		transformer.GetQuery()...,
	)

	return this
}

func (this *mongo_query) SortOrder(fn query.SortFunc) query.ISubQueryBuilder {

	initializer := sort.NewSortInitializer()

	fn(initializer)

	this.mongo_pipeline.Add(
		bson.D{
			{"$sort", initializer},
		},
	)

	return this
}

func (this *mongo_query) Limit(num uint) query.ISubQueryBuilder {

	this.mongo_pipeline.Add(
		bson.D{
			{"$limit", num},
		},
	)

	return this
}
