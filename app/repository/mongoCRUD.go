package repository

import (
	libCommon "app/lib/common"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DESC                  = -1
	ASC                   = 1
	OP_LTE                = "$lte"
	OP_GT                 = "$gt"
	PAGINATION_FIRST_PAGE = 0
	PAGINATION_LAST_PAGE  = 1
)

var (
	empty_bson = bson.D{{}}
)

type (
	Bson_Expression_Type interface {
		bson.D | bson.M
	}
)

func ParseCursor[T any](cursor *mongo.Cursor, ctx context.Context) ([]*T, error) {

	defer cursor.Close(context.TODO())

	var ret []*T

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	for cursor.Next(ctx) {

		var model *T = new(T)

		if err := cursor.Decode(model); err != nil {

			return nil, err
		}

		ret = append(ret, model)
	}

	if err := cursor.Err(); err != nil {

		return nil, err
	}

	return ret, nil
}

func getDocuments[T any](
	page int64,
	collection *mongo.Collection,
	ctx context.Context,
	filters ...interface{},
) ([]*T, error) {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	cursor, err := collection.Aggregate(ctx, bson.D{
		{"$skip", page},
		{"$limit", ITEM_PER_PAGE},
	})

	if err != nil {

		return nil, err
	}

	return ParseCursor[T](cursor, ctx)
}

func getDocumentsPageByID[Model_Type any](
	_id primitive.ObjectID,
	pageLimit int64,
	isPrevDir bool,
	projection *bson.D,
	collection *mongo.Collection,
	ctx context.Context,
	extraFilters ...bson.E,
) ([]*Model_Type, int64, error) {

	if collection == nil {

		panic("no collection provided to retrieve data")
	}

	var paginationQuery bson.D = preparePaginationQuery(_id, isPrevDir, extraFilters)

	fmt.Println(paginationQuery)

	option := options.Find()
	option.Sort = bson.D{{"_id", DESC}}
	option.Limit = &pageLimit

	if projection != nil {

		option.Projection = projection
	}

	cursor, err := collection.Find(ctx, paginationQuery, option)

	if err != nil {

		return nil, 0, err
	}

	data, err := ParseCursor[Model_Type](cursor, context.TODO())

	if err != nil {

		return nil, 0, err
	}

	docCount, err := collection.CountDocuments(ctx, extraFilters)

	if err != nil {

		return nil, 0, err
	}

	return data, docCount, nil
}

func preparePaginationQuery(_id primitive.ObjectID, isPrevDir bool, extraFilters []bson.E) bson.D {

	dir_op := libCommon.Ternary(isPrevDir, OP_LTE, OP_GT)

	//var paginationQuery bson.D

	if _id.IsZero() {

		//paginationQuery = extraFilters

		return bson.D(extraFilters)
	}
	// } else {

	// 	// paginationQuery = bson.D{
	// 	// 	{
	// 	// 		"_id", bson.D{
	// 	// 			{dir_op, _id},
	// 	// 		},
	// 	// 	},
	// 	// }

	// 	// pivotQuery := bson.D{
	// 	// 	{
	// 	// 		"_id", bson.D{
	// 	// 			{dir_op, _id},
	// 	// 		},
	// 	// 	},
	// 	// }

	// 	//paginationQuery = append(pivotQuery, extraFilters...)
	// }

	// if len(extraFilters) > 0 {

	// 	paginationQuery = append(paginationQuery, extraFilters...)
	// }

	//return paginationQuery

	return append(bson.D{
		{
			"_id", bson.D{
				{dir_op, _id},
			},
		},
	}, extraFilters...)
}

func findDocumentByUUID[T any](uuid uuid.UUID, collection *mongo.Collection, ctx context.Context) (*T, error) {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	res := collection.FindOne(ctx, bson.M{
		"uuid": uuid,
	})

	var camp *T

	err := res.Decode(&camp)

	if err != nil {

		return nil, err
	}

	return camp, nil
}

func createDocument[T any](model *T, collection *mongo.Collection, ctx context.Context) error {

	//(*model).UUID = uuid.New()

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	_, err := collection.InsertOne(ctx, model)

	if err != nil {

		return err
	}

	return nil
}

func updateDocument[T any](uuid *uuid.UUID, model *T, collection *mongo.Collection, ctx context.Context) error {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	result, err := collection.UpdateOne(ctx, bson.D{{"uuid", uuid}}, bson.D{{"$set", model}})

	if err != nil {

		return err
	}

	if result.ModifiedCount == 0 {

		return errors.New("No matched document to update")
	}

	return nil
}

func deleteDocument(uuid uuid.UUID, collection *mongo.Collection, ctx context.Context) error {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	_, err := collection.DeleteOne(ctx, bson.D{{"uuid", uuid}})

	if err != nil {

		return err
	}

	return nil
}

func count(collection *mongo.Collection, ctx context.Context, filter ...bson.E) (int64, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	return collection.CountDocuments(ctx, filter)
}
