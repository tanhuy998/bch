package mongoRepository

import (
	"app/internal"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	mongoRepositoryFilter "app/repository/driver/mongod/filter"
	mongoRepositorySorter "app/repository/driver/mongod/sort"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	mongo_repository[Model_T any] struct {
		MongoDBQueryMonitorCollection
		filter     mongoRepositoryFilter.MongoRepositoryFilterGenerator
		sort       mongoRepositorySorter.MongoSorterGenerator
		projection map[string]uint
		//filter     []interface{}
	}
)

func (this *mongo_repository[Model_T]) UpdateOneByUUID(uuid uuid.UUID, model *Model_T, ctx context.Context) error {

	res, err := updateDocument[Model_T](libCommon.PointerPrimitive(uuid), model, &this.MongoDBQueryMonitorCollection, ctx)

	if err != nil {

		return err
	}

	return CheckUpdateOneResult(res)
}

func (this *mongo_repository[Model_T]) UpdateManyByFilter(filter interface{}, update interface{}, ctx context.Context) error {

	res, err := this.MongoDBQueryMonitorCollection.UpdateMany(ctx, filter, bson.D{{"$set", update}})

	if err != nil {

		return err
	}

	return CheckUpdateOneResult(res)
}

func (this *mongo_repository[Model_T]) UpsertManyByFilter(filter interface{}, update interface{}, ctx context.Context) error {

	opts := options.Update().SetUpsert(true)

	_, err := this.MongoDBQueryMonitorCollection.UpdateMany(ctx, filter, bson.D{{"$set", update}}, opts)

	if err != nil {

		return err
	}

	return nil
}

func (this *mongo_repository[Model_T]) DeleteMany(model *Model_T, ctx context.Context) error {

	_, err := this.MongoDBQueryMonitorCollection.DeleteMany(ctx, model)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func (this *mongo_repository[Model_T]) DeleteManyByFilter(filter interface{}, ctx context.Context) error {

	_, err := this.MongoDBQueryMonitorCollection.DeleteMany(ctx, filter)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func (this *mongo_repository[Model_T]) Create(model *Model_T, ctx context.Context) error {

	return createDocument[Model_T](model, &this.MongoDBQueryMonitorCollection, ctx)
}

func (this *mongo_repository[Model_T]) CreateMany(models []*Model_T, ctx context.Context) error {

	_, err := insertMany(models, &this.MongoDBQueryMonitorCollection, context.TODO())

	if err != nil {

		return err
	}

	return nil
}

func (this *mongo_repository[Model_T]) FindByFilter(query bson.D, ctx context.Context) (*Model_T, error) {

	ret, err := findOneDocument[Model_T](query, &this.MongoDBQueryMonitorCollection, ctx, this.projection)

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *mongo_repository[Model_T]) FindManyByFilter(query bson.D, ctx context.Context) ([]*Model_T, error) {

	ret, err := findManyDocuments[Model_T](
		query, &this.MongoDBQueryMonitorCollection, ctx, this.sort.Get(), this.projection,
	)

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *mongo_repository[Model_T]) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*Model_T, error) {

	ret, err := findDocumentByUUID[Model_T](uuid, &this.MongoDBQueryMonitorCollection, ctx)

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *mongo_repository[Model_T]) _FindOffset(
	offset uint64, size uint64, ctx context.Context,
) ([]Model_T, error) {

	findOption := options.Find()
	findOption.Limit = libCommon.PointerPrimitive(int64(size))
	findOption.Sort = this.prepareSorter()
	findOption.Projection = this.projection

	if offset > 1 {

		findOption.Skip = libCommon.PointerPrimitive(int64(offset))
	}

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := this.collection.Find(ctx, this.prepareFilter(), findOption)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	ret, err := ParseValCursor[Model_T](cursor, ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *mongo_repository[Model_T]) FindOffset(
	query bson.D, offset uint64, size uint64, sort *bson.D, ctx context.Context, projection ...bson.E,
) ([]Model_T, error) {

	findOption := options.Find()
	findOption.Limit = libCommon.PointerPrimitive(int64(size))

	if sort == nil {

		findOption.Sort = &bson.D{{"_id", SORT_DESC}}
	} else {

		findOption.Sort = sort
	}

	if offset > 1 {

		findOption.Skip = libCommon.PointerPrimitive(int64(offset))
	}

	if ctx == nil {

		ctx = context.TODO()
	}

	cursor, err := this.collection.Find(ctx, query, findOption)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	ret, err := ParseValCursor[Model_T](cursor, ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *mongo_repository[Model_T]) prepareFilter() []interface{} {

	return this.filter.Get()
}

func (this *mongo_repository[Model_T]) prepareSorter() interface{} {

	return this.sort.Get()
}

func (this *mongo_repository[Model_T]) FindNext(
	cursor primitive.ObjectID, size uint64, ctx context.Context,
) ([]Model_T, error) {

	return FindNext[Model_T](
		this.GetCollection(), internal.PAGINATION_CURSOR_FIELD, cursor, size, ctx, this.prepareFilter(), this.prepareSorter(), this.projection,
	)
}

func (this *mongo_repository[Model_T]) FindPrevious(
	cursor primitive.ObjectID, size uint64, ctx context.Context,
) ([]Model_T, error) {

	return FindPrevious[Model_T](
		this.GetCollection(), internal.PAGINATION_CURSOR_FIELD, cursor, size, ctx, this.prepareFilter(), this.prepareSorter(), this.projection,
	)
}

func (this *mongo_repository[Model_T]) _FindNext(
	cursorField string, cursor interface{}, size uint64, ctx context.Context,
) ([]Model_T, error) {

	return FindNext[Model_T](
		this.GetCollection(), cursorField, cursor, size, ctx, this.prepareFilter(), this.prepareSorter(), this.projection,
	)
}

func (this *mongo_repository[Model_T]) _FindPrevious(
	cursorField string, cursor interface{}, size uint64, ctx context.Context,
) ([]Model_T, error) {

	return FindPrevious[Model_T](
		this.GetCollection(), cursorField, cursor, size, ctx, this.prepareFilter(), this.prepareSorter(), this.projection,
	)
}
