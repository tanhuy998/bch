package mongoRepository

import (
	libCommon "app/internal/lib/common"
	repositoryAPI "app/repository/api"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	mongo_read_projection[Model_T any] struct {
		mongo_filter[Model_T]
	}
)

func (this *mongo_read_projection[Model_T]) InitCollection(col *mongo.Collection) {

	this.collection = col
}

func (this *mongo_read_projection[Model_T]) clone() *mongo_read_projection[Model_T] {

	return libCommon.PointerPrimitive(*this)
}

func (this *mongo_read_projection[Model_T]) Select(fields ...string) (ret repositoryAPI.IRepositoryProjectableOperator[Model_T]) {

	ret = this

	this.initProjection()

	for _, v := range fields {

		if v == "" {

			continue
		}

		this.projection[v] = true
	}

	return
}

func (this *mongo_read_projection[Model_T]) ExcludeFields(fields ...string) (ret repositoryAPI.IRepositoryProjectableOperator[Model_T]) {

	ret = this

	this.initProjection()

	for _, v := range fields {

		if v == "" {

			continue
		}

		this.projection[v] = false
	}

	return
}

func (this *mongo_read_projection[Model_T]) initProjection() {

	if this.projection != nil {

		return
	}

	this.projection = make(map[string]bool)
}

func (this *mongo_read_projection[Model_T]) prepareProjection() bson.D {

	ret := make(bson.D, len(this.projection))

	i := 0

	for field, val := range this.projection {

		ret[i].Key = field
		ret[i].Value = libCommon.Ternary(val, 1, 0)
	}

	return ret
}

func (this *mongo_read_projection[Model_T]) Find(ctx context.Context) ([]*Model_T, error) {

	projection := this.prepareProjection()

	return findManyDocuments[Model_T](this.filter, &this.MongoDBQueryMonitorCollection, ctx, projection...)
}

func (this *mongo_read_projection[Model_T]) FindOne(ctx context.Context) (*Model_T, error) {

	projection := this.prepareProjection()

	return findOneDocument[Model_T](this.filter, &this.MongoDBQueryMonitorCollection, ctx, projection...)
}

func (this *mongo_read_projection[Model_T]) Filter(
	fn repositoryAPI.FilterFunc,
) repositoryAPI.IRepositoryFilterableOperator[Model_T] {

	fn(&this.filter)

	return this
}

// override filter_repository
// func (this *mongo_read_projection[Model_T]) Filter(filter interface{}) repositoryAPI.IRepositoryFilterableOperator[Model_T] {

// 	this.filter = filter

// 	return this
// }
