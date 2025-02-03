package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/lib"
	"app/internal/db/query"
	libCommon "app/internal/lib/common"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	DelegatorQueryBuilder[Model_T any] struct {
		delegator lib.IMongoDBCollection //*MongoDBDelegator[Model_T]
		mongo_query
	}
)

func NewDelegatorQueryBuilder[Model_T any](det lib.IMongoDBCollection) *DelegatorQueryBuilder[Model_T] {

	return &DelegatorQueryBuilder[Model_T]{
		delegator: det,
	}
}

func (this *DelegatorQueryBuilder[Model_T]) Clone() query.IQueryBuilder[Model_T] {

	return libCommon.PointerPrimitive(*this)
}

func (this *DelegatorQueryBuilder[Model_T]) Join(another string, fn func(query.IJoinField)) query.IQueryBuilder[Model_T] {

	this.mongo_query.Join(another, fn)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) Filter(fn query.FilterFunc) query.IQueryBuilder[Model_T] /*named("app/internal/db/query",IFilterableOperator)[tv(Model_T)]*/ {

	this.mongo_query.Filter(fn)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) Select(fields ...string) query.IQueryBuilder[Model_T] {

	this.mongo_query.Select(fields...)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) ExcludeFields(fields ...string) query.IQueryBuilder[Model_T] {

	this.mongo_query.ExcludeFields(fields...)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) Limit() {

}

func (this *DelegatorQueryBuilder[Model_T]) Skip() {

}

func (this *DelegatorQueryBuilder[Model_T]) First(ctx context.Context) (*Model_T, error) {

	this.mongo_pipeline.Add(
		bson.D{
			{"$limit", 1},
		},
	)

	return lib.AggregateOne[Model_T](
		this.delegator,
		&this.mongo_pipeline,
		ctx,
	)
}

func (this *DelegatorQueryBuilder[Model_T]) All(ctx context.Context) ([]Model_T, error) {

	return lib.Aggregate[Model_T](
		this.delegator,
		&this.mongo_pipeline,
		ctx,
	)
}

func (this *DelegatorQueryBuilder[Model_T]) Transform(fn query.DataTransformFunc) query.IQueryBuilder[Model_T] {

	this.mongo_query.Transform(fn)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) NewExecutor(delegator query.IDBDelegator[Model_T]) query.IQueryExecutor[Model_T] {

	d, ok := delegator.(lib.IMongoDBCollection)

	if !ok {

		panic("invalid concrete value of interface query.IDBDelegator")
	}

	ret := libCommon.PointerPrimitive(*this)

	ret.delegator = d

	return ret
}

func (d *DelegatorQueryBuilder[Model_T]) SortOrder(fn query.SortFunc) query.IQueryBuilder[Model_T] {
	panic("TODO: Implement")
}
