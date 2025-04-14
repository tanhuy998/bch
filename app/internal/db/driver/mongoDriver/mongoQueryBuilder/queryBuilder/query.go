package queryBuilder

import (
	"app/internal/db/driver/mongoDriver/filter"
	"app/internal/db/driver/mongoDriver/lib"
	libCommon "app/internal/lib/common"
	"fmt"

	"app/internal/db/driver/mongoDriver/sort"
	"app/internal/db/driver/mongoDriver/transform"
	"app/internal/db/query"
	"app/internal/db/storage"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	MongoAggregateQueryBuilder struct {
		lib.MongoPipeline
		projection map[string]uint
	}
)

func (this *MongoAggregateQueryBuilder) GetPipeline() []interface{} {

	return this.MongoPipeline.P
}

func (this *MongoAggregateQueryBuilder) initProjection() {

	if this.projection != nil {

		return
	}

	this.projection = make(map[string]uint)
}

func (this *MongoAggregateQueryBuilder) Join(
	another storage.IDBStorageIdentifier, fn func(query.IJoinField),
) query.IJoinUnwindableQueryBuilder /* query.IQueryBuilder */ {

	joinIntializer := NewJoinOperationQueryBuilder()
	joinIntializer.From = another.GetDBStorageUnitName()

	fn(joinIntializer)

	joinIntializer.Done()

	this.MongoPipeline.PushStages(
		bson.D{
			{"$lookup", joinIntializer},
		},
	)

	//return this

	return NewJoinUnwindableDelegator(
		libCommon.Ternary(joinIntializer.Alias == "", fmt.Sprintf(`%ss`, another.GetDBStorageUnitName()), joinIntializer.Alias),
		this,
	)
}

func (this *MongoAggregateQueryBuilder) Filter(fn query.FilterFunc) query.IQueryBuilder {

	filter := filter.NewFilterGenerator()

	fn(filter)

	this.MongoPipeline.PushStages(
		bson.D{
			{"$match", filter.Get()},
		},
	)

	return this
}

func (this *MongoAggregateQueryBuilder) Select(fields ...string) query.IQueryBuilder {

	this.initProjection()

	for _, name := range fields {

		this.projection[name] = 1
	}

	return this
}

func (this *MongoAggregateQueryBuilder) ExcludeFields(fields ...string) query.IQueryBuilder {

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

func (this *MongoAggregateQueryBuilder) Transform(fn query.DataTransformFunc) query.IQueryBuilder {

	transformer := transform.NewDataTransformer()

	fn(transformer)

	this.MongoPipeline.PushStages(
		transformer.GetQuery()...,
	)

	return this
}

func (this *MongoAggregateQueryBuilder) SortOrder(fn query.SortFunc) query.IQueryBuilder {

	initializer := sort.NewSortInitializer()

	fn(initializer)

	this.MongoPipeline.PushStages(
		bson.D{
			{"$sort", initializer.GetMap()},
		},
	)

	return this
}

func (this *MongoAggregateQueryBuilder) Limit(num uint64) query.IQueryBuilder {

	this.MongoPipeline.PushStages(
		bson.D{
			{"$limit", num},
		},
	)

	return this
}

func (this *MongoAggregateQueryBuilder) Clone() query.IClonableQueryBuilder {

	return this._clone()
}

func (this *MongoAggregateQueryBuilder) CloneThis() *MongoAggregateQueryBuilder {

	return this._clone()
}

func (this *MongoAggregateQueryBuilder) _clone() *MongoAggregateQueryBuilder {

	ret := new(MongoAggregateQueryBuilder)

	ret.P = make([]interface{}, len(this.P))
	copy(ret.P, this.P)

	return ret
}
func (this *MongoAggregateQueryBuilder) Skip(time uint64) query.IQueryBuilder {

	this.MongoPipeline.PushStages(
		bson.D{
			{"$skip", time},
		},
	)

	return this
}

func (this *MongoAggregateQueryBuilder) Unwind(field string) {

	this.MongoPipeline.PushStages(
		bson.D{
			{
				"$unwind", bson.D{
					{"path", field},
					//{"includeArrayIndex", false},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
	)
}
