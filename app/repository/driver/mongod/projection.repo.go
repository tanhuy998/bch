package mongoRepository

import (
	repositoryAPI "app/repository/api"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	mongo_read_projection[Model_T any] struct {
		//mongo_filter[Model_T]
		mongo_repository[Model_T]
	}
)

func (this *mongo_read_projection[Model_T]) InitCollection(col *mongo.Collection) {

	this.collection = col
}

func (this *mongo_read_projection[Model_T]) Select(fields ...string) (ret repositoryAPI.IRepositoryProjectableOperator[Model_T]) {

	ret = this

	this.initProjection()

	for _, v := range fields {

		if v == "" {

			continue
		}

		this.projection[v] = 1
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

		this.projection[v] = 0
	}

	return
}

func (this *mongo_read_projection[Model_T]) initProjection() {

	if this.projection != nil {

		return
	}

	this.projection = make(map[string]uint)
}

func (this *mongo_read_projection[Model_T]) Find(ctx context.Context) ([]*Model_T, error) {

	return findManyDocuments[Model_T](this.prepareFilter(), &this.MongoDBQueryMonitorCollection, ctx, this.prepareSorter(), this.projection)
}

func (this *mongo_read_projection[Model_T]) FindOne(ctx context.Context) (*Model_T, error) {

	return findOneDocument[Model_T](this.prepareFilter(), &this.MongoDBQueryMonitorCollection, ctx, this.projection)
}

func (this *mongo_read_projection[Model_T]) FindNext(
	cursorField string, cursor interface{}, size uint64, ctx context.Context,
) ([]Model_T, error) {

	return this._FindNext(
		cursorField, cursor, size, ctx,
	)
}

func (this *mongo_read_projection[Model_T]) FindPrevious(
	cursorField string, cursor interface{}, size uint64, ctx context.Context,
) ([]Model_T, error) {

	return this._FindPrevious(
		cursorField, cursor, size, ctx,
	)
}

func (this *mongo_read_projection[Model_T]) FindOffset(
	offset uint64, size uint64, ctx context.Context,
) ([]Model_T, error) {

	return this.mongo_repository._FindOffset(
		offset, size, ctx,
	)
}
