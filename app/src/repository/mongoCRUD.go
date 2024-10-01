package repository

import (
	libCommon "app/src/internal/lib/common"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	Bson_Expression_Type interface {
		bson.D | bson.M
	}

	MongoDBCursorSortOrder = int
)

const (
	SORT_DESC             MongoDBCursorSortOrder = -1
	SORT_ASC              MongoDBCursorSortOrder = 1
	OP_LTE                                       = "$lte"
	OP_LT                                        = "$lt"
	OP_GT                                        = "$gt"
	OP_GTE                                       = "$gte"
	PAGINATION_FIRST_PAGE                        = 0
	PAGINATION_LAST_PAGE                         = 1
)

var (
	empty_bson = bson.D{{}}
)

func insertMany[T any](
	models []*T,
	collection *mongo.Collection,
	ctx context.Context,
) (*mongo.InsertManyResult, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var documents []interface{} = make([]interface{}, len(models))

	for i, model := range models {

		documents[i] = model
	}

	return collection.InsertMany(
		ctx, documents,
	)
}

func findManyDocuments[T any](
	query bson.D,
	collection *mongo.Collection,
	ctx context.Context,
	projections ...bson.E,
) ([]*T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var opts *options.FindOptions

	if len(projections) > 0 {

		opts = options.Find().SetProjection(projections)
	}

	cur, err := collection.Find(ctx, query, opts)

	if err != nil {

		return nil, err
	}

	res, err := ParseCursor[T](cur, ctx)

	if err != nil {

		return nil, err
	}

	return res, nil
}

func ParseCursor[T any](cursor *mongo.Cursor, ctx context.Context) ([]*T, error) {

	defer cursor.Close(context.TODO())

	var ret []*T

	if ctx == nil {

		ctx = context.TODO()
	}

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

func findOneDocument[T any](
	query bson.D,
	collection *mongo.Collection,
	ctx context.Context,
	projections ...bson.E,
) (*T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var model *T = new(T)

	opts := options.FindOne()

	if len(projections) > 0 {

		opts.Projection = projections
	}

	err := collection.FindOne(ctx, query, opts).Decode(model)

	if err != nil {

		return nil, err
	}

	return model, nil
}

func getDocuments[T any](
	page int64,
	collection *mongo.Collection,
	ctx context.Context,
	filters ...interface{},
) ([]*T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

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
) (*PaginationPack[Model_Type], error) {

	if collection == nil {

		panic("no collection provided to retrieve data")
	}

	session, err := collection.Database().Client().StartSession()

	if err != nil {

		return nil, err
	}
	defer session.EndSession(ctx)

	writeConcern := writeconcern.Majority()
	readConcern := readconcern.Snapshot()
	transactionOpts := options.Transaction().SetWriteConcern(writeConcern).SetReadConcern(readConcern)

	pack, err := session.WithTransaction(
		ctx,
		func(mongo.SessionContext) (interface{}, error) {

			var paginationQuery bson.D = PrepareObjIDFilterPaginationQuery(_id, isPrevDir, extraFilters)
			var sortOrder MongoDBCursorSortOrder = SORT_DESC

			if isPrevDir {

				sortOrder = SORT_ASC
			}

			option := options.Find()
			// option.Sort = bson.D{{"_id", SORT_DESC}}
			option.Sort = bson.D{{"_id", sortOrder}}
			option.Limit = &pageLimit

			if projection != nil {

				option.Projection = projection
			}

			cursor, err := collection.Find(ctx, paginationQuery, option)

			if err != nil {

				return nil, err
			}

			data, err := ParseCursor[Model_Type](cursor, context.TODO())

			if err != nil {

				return nil, err
			}

			var filters bson.D

			if len(extraFilters) == 0 {

				filters = empty_bson
			} else {

				filters = extraFilters
			}

			docCount, err := collection.CountDocuments(ctx, filters)

			if err != nil {

				return nil, err
			}

			if isPrevDir {

				data = libCommon.ReverseSlice(data...)
			}

			dataPack := PaginationPack[Model_Type]{
				Data:  data,
				Count: docCount,
			}

			return dataPack, nil
		},
		transactionOpts,
	)

	if err != nil {

		return nil, err
	}

	if packActualVal, ok := pack.(PaginationPack[Model_Type]); ok {

		return &packActualVal, nil
	}

	return nil, errors.New("error while unpacking pagination data")
}

func initDBTransaction(client *mongo.Client) (*mongo.Session, error) {

	session, err := client.StartSession()

	if err != nil {

		return nil, err
	}

	return &session, nil
}

func PrepareObjIDFilterPaginationQuery(_id primitive.ObjectID, isPrevDir bool, extraFilters []bson.E) bson.D {

	var dir_op string

	if isPrevDir {

		dir_op = OP_GT
	} else {

		dir_op = OP_LT
	}

	if _id.IsZero() {

		//paginationQuery = extraFilters

		if len(extraFilters) == 0 {

			return empty_bson
		}

		return bson.D(extraFilters)
	}

	return append(bson.D{
		{
			"_id", bson.D{
				{dir_op, _id},
			},
		},
	}, extraFilters...)
}

func PrepareAggregatePaginationQuery(paginationPivotField string, pivotValue interface{}, isPrevDir bool, extraFilters []bson.E) bson.D {

	var dir_op string

	if isPrevDir {

		dir_op = OP_GT
	} else {

		dir_op = OP_LT
	}

	if pivotValue == nil {

		dir_op = OP_GTE
	}

	return append(bson.D{
		{
			paginationPivotField, bson.D{
				{dir_op, pivotValue},
			},
		},
	}, extraFilters...)
}

func findDocumentByUUID[T any](uuid uuid.UUID, collection *mongo.Collection, ctx context.Context, projections ...bson.E) (*T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	opts := options.FindOne()

	if len(projections) > 0 {

		opts.Projection = projections
	}

	res := collection.FindOne(
		ctx,
		bson.M{
			"uuid": uuid,
		},
		opts,
	)

	var camp *T

	err := res.Decode(&camp)

	if err != nil {

		return nil, err
	}

	return camp, nil
}

func createDocument[T any](model *T, collection *mongo.Collection, ctx context.Context) error {

	//(*model).UUID = uuid.New()

	if ctx == nil {

		ctx = context.TODO()
	}

	_, err := collection.InsertOne(ctx, model)

	if err != nil {

		return err
	}

	return nil
}

func updateDocument[T any](uuid *uuid.UUID, model *T, collection *mongo.Collection, ctx context.Context, extraFilters ...bson.E) (*mongo.UpdateResult, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var targetFilter bson.D = bson.D{{"uuid", uuid}}

	if len(extraFilters) > 0 {

		targetFilter = append(targetFilter, extraFilters...)
	}

	result, err := collection.UpdateOne(ctx, targetFilter, bson.D{{"$set", model}})

	if err != nil {

		return nil, err
	}

	return result, nil
}

func UpdateOneByUUID[T any](uuid *uuid.UUID, model *T, collection *mongo.Collection, ctx context.Context, extraFilters ...bson.E) (*mongo.UpdateResult, error) {

	return updateDocument(uuid, model, collection, ctx, extraFilters...)
}

func deleteDocument(uuid uuid.UUID, collection *mongo.Collection, ctx context.Context) error {

	if ctx == nil {

		ctx = context.TODO()
	}

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

func Aggregate[Model_T any](collection *mongo.Collection, pipeline mongo.Pipeline, ctx context.Context, options ...*options.AggregateOptions) ([]*Model_T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := collection.Aggregate(ctx, pipeline, options...)

	if err != nil {

		return nil, err
	}

	return ParseCursor[Model_T](cursor, context.TODO())
}

func AggregateByPage[Model_T any](
	collection *mongo.Collection,
	pipeline mongo.Pipeline,
	paginationPivotField string,
	pivotValue interface{},
	pageLimit int64,
	isPrevDir bool,
	pipelineAfterPivot mongo.Pipeline,
	ctx context.Context,
	option ...*options.AggregateOptions,
) (*PaginationPack[Model_T], error) {

	if collection == nil {

		panic("no collection provided for aggregation")
	}

	if ctx == nil {

		ctx = context.TODO()
	}

	paginationStages := prepareAggregationPaginationStages(paginationPivotField, pivotValue, isPrevDir)
	pipeline = append(pipeline, paginationStages...)

	if pipelineAfterPivot != nil {

		pipeline = append(pipeline, pipelineAfterPivot...)
	}

	resData, err := Aggregate[Model_T](collection, pipeline, ctx)

	if err != nil {

		return nil, err
	}

	dataPack := &PaginationPack[Model_T]{
		Data: resData,
	}

	return dataPack, nil
}

func prepareAggregationPaginationStages(
	paginationPivotField string,
	pivotValue interface{},
	isPrevDir bool,
) mongo.Pipeline {

	if paginationPivotField == "" {

		paginationPivotField = "_id"
	}

	pivotQuery := PrepareAggregatePaginationQuery(paginationPivotField, pivotValue, isPrevDir, nil)
	pivotStage := bson.D{
		{"$match", pivotQuery},
	}

	var sortOrder MongoDBCursorSortOrder = SORT_DESC

	if isPrevDir {

		sortOrder = SORT_ASC
	}

	sortStage := bson.D{
		{
			"$sort", bson.D{
				{paginationPivotField, sortOrder},
			},
		},
	}

	return mongo.Pipeline{
		pivotStage, sortStage,
	}
}
