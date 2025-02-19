package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/lib"
	"app/internal/db/query"
	"app/internal/db/storage"
	libCommon "app/internal/lib/common"
)

type (
	DelegatorQueryBuilder[Model_T any] struct {
		// delegator lib.IMongoDBCollection //*MongoDBDelegator[Model_T]
		// mongo_query
		//QueryExecutor[Model_T]
		PaginateExecutor[Model_T]
	}
)

func NewDelegatorQueryBuilder[Model_T any](det lib.IMongoDBCollection) *DelegatorQueryBuilder[Model_T] {

	// return &DelegatorQueryBuilder[Model_T]{
	// 	delegator: det,
	// }

	ret := new(DelegatorQueryBuilder[Model_T])

	ret.delegator = det

	return ret
}

func (this *DelegatorQueryBuilder[Model_T]) Clone() query.IQueryBuilder[Model_T] {

	return libCommon.PointerPrimitive(*this)
}

func (this *DelegatorQueryBuilder[Model_T]) Join(storage storage.IDBStorageIdentifier, fn func(query.IJoinField)) query.IQueryBuilder[Model_T] {

	this.MongoAggregateQueryBuilder.Join(storage.GetDBStorageUnitName(), fn)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) Filter(fn query.FilterFunc) query.IQueryBuilder[Model_T] /*named("app/internal/db/query",IFilterableOperator)[tv(Model_T)]*/ {

	this.MongoAggregateQueryBuilder.Filter(fn)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) Select(fields ...string) query.IQueryBuilder[Model_T] {

	this.MongoAggregateQueryBuilder.Select(fields...)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) ExcludeFields(fields ...string) query.IQueryBuilder[Model_T] {

	this.MongoAggregateQueryBuilder.ExcludeFields(fields...)

	return this
}

func (this *DelegatorQueryBuilder[Model_T]) Limit() {

}

func (this *DelegatorQueryBuilder[Model_T]) Skip() {

}

func (this *DelegatorQueryBuilder[Model_T]) Transform(fn query.DataTransformFunc) query.IQueryBuilder[Model_T] {

	this.MongoAggregateQueryBuilder.Transform(fn)

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

func (this *DelegatorQueryBuilder[Model_T]) SortOrder(fn query.SortFunc) query.IQueryBuilder[Model_T] {
	panic("TODO: Implement")
}

// func (this *DelegatorQueryBuilder[Model_T]) Paginate(p0 query.PaginateInitFunc[interface{}]) ([]Model_T, error) {
// 	panic("TODO: Implement")
// }
