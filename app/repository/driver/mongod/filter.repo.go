package mongoRepository

import (
	"app/internal/common"
	repositoryAPI "app/repository/api"
	"context"
	"errors"
	"fmt"
)

type (
	mongo_filter[Model_T any] struct {
		//mongo_repository[Model_T]
		mongo_read_projection[Model_T]
	}
)

func (this *mongo_filter[Model_T]) Update(updateEntity Model_T, ctx context.Context) error {

	return this.UpdateManyByFilter(this.prepareFilter(), updateEntity, ctx)
}

func (this *mongo_filter[Model_T]) UpdateOne(updateEntity Model_T, ctx context.Context) error {

	if ctx == nil {

		ctx = context.TODO()
	}

	updateRes, err := this.MongoDBQueryMonitorCollection.UpdateOne(
		ctx, this.filter, updateEntity,
	)

	if err != nil {

		return err
	}

	if updateRes.MatchedCount == 0 {

		return errors.Join(
			common.ERR_NOT_FOUND,
			fmt.Errorf("update on inexisting resource"),
		)
	}

	if updateRes.ModifiedCount == 0 {

		return errors.Join(
			common.ERR_CONFLICT,
			fmt.Errorf("duplication of resource update"),
		)
	}

	return nil
}

func (this *mongo_filter[Model_T]) Delete(ctx context.Context) error {

	return this.DeleteManyByFilter(this.filter, ctx)
}

func (this *mongo_filter[Model_T]) DeleteOne(ctx context.Context) error {

	if ctx == nil {

		ctx = context.TODO()
	}

	_, err := this.MongoDBQueryMonitorCollection.DeleteOne(ctx, this.filter)

	if err != nil {

		return err
	}

	return nil
}

func (this *mongo_filter[Model_T]) Upsert(entity Model_T, ctx context.Context) error {

	return this.UpsertManyByFilter(this.filter, entity, ctx)
}

// func (this *mongo_filter[Model_T]) FindNext(cursor interface{}, size uint64, ctx context.Context) ([]Model_T, error) {

// 	return this._FindNext(internal.PAGINATION_CURSOR_FIELD, cursor, size, ctx)
// }

// func (this *mongo_filter[Model_T]) FindPrevious(cursor interface{}, size uint64, ctx context.Context) ([]Model_T, error) {

// 	return this._FindPrevious(internal.PAGINATION_CURSOR_FIELD, cursor, size, ctx)
// }

func (this *mongo_filter[Model_T]) Filter(
	fn repositoryAPI.FilterFunc,
) repositoryAPI.IRepositoryFilterableOperator[Model_T] {

	fn(&this.filter)

	return this
}
