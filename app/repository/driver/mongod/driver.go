package mongoRepository

import (
	"app/internal"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"context"
	"errors"
	"slices"

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

	/*
		Models that inplement this interface must define
		two mehod to return the pagination query based on
		current value of the model
	*/
	ICustomCursorPaginationModel interface {
		QueryBefore() bson.D
		QueryAfter() bson.D
	}
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
	collection IMongoRepositoryOperator,
	ctx context.Context,
) (*mongo.InsertManyResult, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var documents []interface{} = make([]interface{}, len(models))

	for i, model := range models {

		documents[i] = model
	}

	ret, err := collection.InsertMany(
		ctx, documents,
	)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ret, nil
}

func findManyDocuments[T any](
	query interface{},
	collection IMongoRepositoryOperator,
	ctx context.Context,
	sort interface{},
	//projections ...bson.E,
	projection interface{},
) ([]*T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var opts *options.FindOptions

	// if len(projections) > 0 {

	// 	opts = options.Find().SetProjection(projections)
	// }

	opts.Projection = projection
	opts.Sort = sort

	cur, err := collection.Find(ctx, query, opts)

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, libError.NewInternal(err)
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

func findOneDocument[T any](
	query interface{},
	collection IMongoRepositoryOperator,
	ctx context.Context,
	//projections ...bson.E,
	projection interface{},
) (*T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var model *T = new(T)

	opts := options.FindOne()

	// if len(projections) > 0 {

	// 	opts.Projection = projections
	// }

	opts.Projection = projection

	res := collection.FindOne(ctx, query, opts)

	err := res.Err()

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	res.Decode(model)

	return model, nil
}

func getDocuments[T any](
	page int64,
	collection IMongoRepositoryOperator,
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

		return nil, libError.NewInternal(err)
	}

	return ParseCursor[T](cursor, ctx)
}

func getDocumentsPageByID[Model_Type any](
	_id primitive.ObjectID,
	pageLimit int64,
	isPrevDir bool,
	projection *bson.D,
	collection IMongoRepositoryOperator,
	ctx context.Context,
	extraFilters ...interface{},
) (*PaginationPack[Model_Type], error) {

	if collection == nil {

		panic("no collection provided to retrieve data")
	}

	session, err := collection.Database().Client().StartSession()

	if err != nil {

		return nil, libError.NewInternal(err)
	}
	defer session.EndSession(ctx)

	writeConcern := writeconcern.Majority()
	readConcern := readconcern.Snapshot()
	transactionOpts := options.Transaction().SetWriteConcern(writeConcern).SetReadConcern(readConcern)

	pack, err := session.WithTransaction(
		ctx,
		func(mongo.SessionContext) (interface{}, error) {

			var paginationQuery interface{} = PrepareObjIDFilterPaginationQuery(_id, isPrevDir, extraFilters)
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

				return nil, libError.NewInternal(err)
			}

			data, err := ParseCursor[Model_Type](cursor, context.TODO())

			if err != nil {

				return nil, err
			}

			var filters interface{}

			if len(extraFilters) == 0 {

				filters = empty_bson
			} else {

				filters = extraFilters
			}

			docCount, err := collection.CountDocuments(ctx, filters)

			if err != nil {

				return nil, libError.NewInternal(err)
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

	return nil, libError.NewInternal(errors.New("error while unpacking pagination data"))
}

func initDBTransaction(client *mongo.Client) (*mongo.Session, error) {

	session, err := client.StartSession()

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return &session, nil
}

func PrepareObjIDFilterPaginationQuery(_id primitive.ObjectID, isPrevDir bool, extraFilters []interface{}) interface{} {

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

		return []interface{}{extraFilters} //bson.D(extraFilters)
	}

	return append([]interface{}{
		bson.E{
			"_id", bson.D{
				{dir_op, _id},
			},
		},
	}, extraFilters...)

	// return append(bson.D{
	// 	{
	// 		"_id", bson.D{
	// 			{dir_op, _id},
	// 		},
	// 	},
	// }, extraFilters...)
}

func PrepareAggregatePaginationQuery(paginationPivotField string, pivotValue interface{}, isPrevDir bool, extraFilters []interface{}) interface{} {

	var dir_op string

	if isPrevDir {

		dir_op = OP_GT
	} else {

		dir_op = OP_LT
	}

	if pivotValue == nil {

		dir_op = OP_GTE
	}

	return append(
		[]interface{}{
			bson.E{
				paginationPivotField, bson.D{
					{dir_op, pivotValue},
				},
			},
		},
		extraFilters...,
	)

	// return append(bson.D{
	// 	{
	// 		paginationPivotField, bson.D{
	// 			{dir_op, pivotValue},
	// 		},
	// 	},
	// }, extraFilters...)
}

func findDocumentByUUID[T any](uuid uuid.UUID, collection IMongoRepositoryOperator, ctx context.Context, projections ...bson.E) (*T, error) {

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

		return nil, libError.NewInternal(err)
	}

	return camp, nil
}

func createDocument[T any](model *T, collection IMongoRepositoryOperator, ctx context.Context) error {

	//(*model).UUID = uuid.New()

	if ctx == nil {

		ctx = context.TODO()
	}

	_, err := collection.InsertOne(ctx, model)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func updateDocument[T any](uuid *uuid.UUID, model *T, collection IMongoRepositoryOperator, ctx context.Context, extraFilters ...bson.E) (*mongo.UpdateResult, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	var targetFilter bson.D = bson.D{{"uuid", uuid}}

	if len(extraFilters) > 0 {

		targetFilter = append(targetFilter, extraFilters...)
	}

	result, err := collection.UpdateOne(ctx, targetFilter, bson.D{{"$set", model}})

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return result, nil
}

func UpdateOneByUUID[T any](uuid *uuid.UUID, model *T, collection IMongoRepositoryOperator, ctx context.Context, extraFilters ...bson.E) (*mongo.UpdateResult, error) {

	return updateDocument(uuid, model, collection, ctx, extraFilters...)
}

func deleteDocument(uuid uuid.UUID, collection IMongoRepositoryOperator, ctx context.Context) error {

	if ctx == nil {

		ctx = context.TODO()
	}

	_, err := collection.DeleteOne(ctx, bson.D{{"uuid", uuid}})

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func count(collection IMongoRepositoryOperator, ctx context.Context, filter ...bson.E) (int64, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	return collection.CountDocuments(ctx, filter)
}

func Aggregate[Model_T any](
	collection IMongoRepositoryOperator, pipeline mongo.Pipeline, ctx context.Context, options ...*options.AggregateOptions,
) ([]*Model_T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := collection.Aggregate(ctx, pipeline, options...)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ParseCursor[Model_T](cursor, context.TODO())
}

/*
return the first document of the aggregated result
*/
func AggregateOne[Model_T any](
	collection IMongoRepositoryOperator, pipeline mongo.Pipeline, ctx context.Context, options ...*options.AggregateOptions,
) (*Model_T, error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := collection.Aggregate(ctx, pipeline, options...)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ParseCursorOne[Model_T](cursor, ctx)
}

// func AgggregateCursor[Model_T any]([]Model_T, error) (
// 	collection IMongoRepositoryOperator,
// ) {

// }

func AggregateByPage[Model_T any](
	collection IMongoRepositoryOperator,
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

		return nil, libError.NewInternal(err)
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

// func FindNext[Model_T libMongo.IBsonDocument](
// 	collection IMongoRepositoryOperator, dataModel Model_T, size uint64, ctx context.Context, filters ...bson.E,
// ) ([]Model_T, error) {

// 	if size == 0 {

// 		size = DEFAULT_PAGINATION_SIZE
// 	}

// 	query := make(bson.D, len(filters)+1)

// 	query[0] = bson.E{
// 		"_id", bson.D{
// 			{"$lt", dataModel.GetObjectID()},
// 		},
// 	}

// 	for i, f := range filters {

// 		query[i+1] = f
// 	}

// 	if ctx == nil {

// 		ctx = context.TODO()
// 	}

// 	findOption := options.Find()
// 	findOption.Limit = libCommon.PointerPrimitive(int64(size))
// 	findOption.Sort = bson.D{{"_id", SORT_DESC}}

// 	c, err := collection.Find(ctx, query, findOption)

// 	if err != nil {

// 		return nil, libError.NewInternal(err)
// 	}

// 	return ParseValCursor[Model_T](c, ctx)
// }

// func FindPrevious[Model_T libMongo.IBsonDocument](
// 	collection IMongoRepositoryOperator, dataModel Model_T, size uint64, ctx context.Context, filters ...bson.E,
// ) ([]Model_T, error) {

// 	if size == 0 {

// 		size = DEFAULT_PAGINATION_SIZE
// 	}

// 	query := make(bson.D, len(filters)+1)

// 	query[0] = bson.E{
// 		"_id", bson.D{
// 			{"$gt", dataModel.GetObjectID()},
// 		},
// 	}

// 	for i, f := range filters {

// 		query[i+1] = f
// 	}

// 	if ctx == nil {

// 		ctx = context.TODO()
// 	}

// 	findOption := options.Find()
// 	findOption.Limit = libCommon.PointerPrimitive(int64(size))
// 	findOption.Sort = bson.D{{"_id", SORT_ASC}}

// 	c, err := collection.Find(ctx, query, findOption)

// 	if err != nil {

// 		return nil, libError.NewInternal(err)
// 	}

// 	ret, err := ParseValCursor[Model_T](c, ctx)

// 	if err != nil {

// 		return nil, err
// 	}

// 	slices.Reverse(ret)

// 	return ret, nil
// }

func FindNext[Entity_T any, Cursor_T comparable, Filter_T any](
	collection IMongoRepositoryOperator, cursorField string, cursor Cursor_T, size uint64, ctx context.Context, filters []Filter_T, sort interface{}, projection interface{},
) ([]Entity_T, error) {

	if cursorField == "" {

		cursorField = internal.PAGINATION_CURSOR_FIELD
	}

	if size == 0 {

		size = DEFAULT_PAGINATION_SIZE
	}

	query := make([]interface{}, len(filters)+1)

	query[0] = bson.E{
		cursorField, bson.D{
			{"$lt", cursor},
		},
	}

	for i, f := range filters {

		query[i+1] = f
	}

	if ctx == nil {

		ctx = context.TODO()
	}

	findOption := options.Find()
	findOption.Limit = libCommon.PointerPrimitive(int64(size))
	findOption.Sort = bson.D{{cursorField, SORT_DESC}}
	findOption.Projection = projection

	c, err := collection.Find(ctx, query, findOption)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return ParseValCursor[Entity_T](c, ctx)
}

func FindPrevious[Entity_T any, Cursor_T comparable, Filter_T any](
	collection IMongoRepositoryOperator, cursorField string, cursor Cursor_T, size uint64, ctx context.Context, filters []Filter_T, sort interface{}, projection interface{},
) ([]Entity_T, error) {

	if cursorField == "" {

		cursorField = internal.PAGINATION_CURSOR_FIELD
	}

	if size == 0 {

		size = DEFAULT_PAGINATION_SIZE
	}

	query := make([]interface{}, len(filters)+1)

	query[0] = bson.E{
		cursorField, bson.D{
			{"$gt", cursor},
		},
	}

	for i, f := range filters {

		query[i+1] = f
	}

	if ctx == nil {

		ctx = context.TODO()
	}

	findOption := options.Find()
	findOption.Limit = libCommon.PointerPrimitive(int64(size))
	findOption.Sort = bson.D{{cursorField, SORT_ASC}}
	findOption.Projection = projection

	c, err := collection.Find(ctx, query, findOption)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	ret, err := ParseValCursor[Entity_T](c, ctx)

	if err != nil {

		return nil, err
	}

	slices.Reverse(ret)

	return ret, nil
}
