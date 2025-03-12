package lib

import (
	libError "app/internal/lib/error"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ParseCursor[T any](cursor *mongo.Cursor, ctx context.Context) ([]*T, error) {

	defer cursor.Close(context.TODO())

	var ret []*T

	if ctx == nil {

		ctx = context.TODO()
	}

	for cursor.Next(ctx) {

		var model *T = new(T)

		if err := cursor.Decode(model); err != nil {

			return nil, libError.NewInternal(err)
		}

		ret = append(ret, model)
	}

	if err := cursor.Err(); err != nil {

		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func ParseValCursor[T any](cursor *mongo.Cursor, ctx context.Context) ([]T, error) {

	ret := make([]T, 0)

	err := cursor.All(ctx, &ret)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func ParseCursorOne[T any](cursor *mongo.Cursor, ctx context.Context) (*T, error) {

	defer cursor.Close(context.TODO())

	if ctx == nil {

		ctx = context.TODO()
	}

	hasDocument := cursor.Next(ctx)

	if cursor.Err() != nil {

		return nil, libError.NewInternal(cursor.Err())
	}

	if !hasDocument {

		return nil, nil
	}

	var ret *T = new(T)

	err := cursor.Decode(ret)

	if err != nil {
		//fmt.Println(reflect.TypeOf(err))
		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func AggregateRaw(
	result interface{}, collection IMongoDBCollection, pipeline interface{}, ctx context.Context, options ...*options.AggregateOptions,
) error {

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := collection.Aggregate(ctx, pipeline, options...)

	if err != nil {

		return libError.NewInternal(
			errors.New("mongo driver aggregate raw error while marshalling query:"),
			err,
		)
	}

	err = cursor.All(ctx, result)

	if err != nil {

		return libError.NewInternal(
			errors.New("mongo driver aggregate raw error while decoding result:"),
			err,
		)
	}

	return nil
}

func AggregateRawOne(
	result interface{}, collection IMongoDBCollection, pipeline interface{}, ctx context.Context, options ...*options.AggregateOptions,
) error {

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := collection.Aggregate(ctx, pipeline, options...)

	if err != nil {

		return libError.NewInternal(
			errors.New("mongo driver aggregate raw one error while marshalling query:"),
			err,
		)
	}

	hasDocument := cursor.Next(ctx)

	if !hasDocument {

		result = nil
		return nil
	}

	err = cursor.Decode(result)

	if err != nil {

		return libError.NewInternal(
			errors.New("mongo driver aggregate raw one error while decoding result:"),
			err,
		)
	}

	return nil
}

func Aggregate[Model_T any](
	collection IMongoDBCollection, pipeline interface{}, ctx context.Context, options ...*options.AggregateOptions,
) ([]Model_T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := collection.Aggregate(ctx, pipeline, options...)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ParseValCursor[Model_T](cursor, context.TODO())
}

func AggregateOne[Model_T any](
	collection IMongoDBCollection, pipeline interface{}, ctx context.Context, options ...*options.AggregateOptions,
) (*Model_T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := collection.Aggregate(ctx, pipeline, options...)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ParseCursorOne[Model_T](cursor, context.TODO())
}
