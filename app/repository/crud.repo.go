package repository

import (
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DEFAULT_PAGINATION_SIZE = 10
)

type (
	IFindMany[Model_T any] interface {
		FindMany(query bson.D, ctx context.Context, projection ...bson.E) ([]*Model_T, error)
	}

	ICreateMany[Model_T any] interface {
		CreateMany(models []*Model_T, ctx context.Context) error
	}

	IFindByFilterRepository[Model_T any] interface {
		Find(filter interface{}, ctx context.Context) (*Model_T, error)
		FindOffset(
			filter interface{}, offset uint64, size uint64, sort *bson.D, ctx context.Context, projection ...bson.E,
		) ([]Model_T, error)
		FindMany(filter interface{}, ctx context.Context, projection ...bson.E) ([]*Model_T, error)
	}

	IUpdateByFilterRepository[Model_T any] interface {
		UpdateManyByFilter(filter interface{}, update bson.D, ctx context.Context) error
		UpsertManyByFilter(filter interface{}, update bson.D, ctx context.Context) error
	}

	IDeleteByFilterRepository[Model_T any] interface {
		DeleteManyByFilter(filter interface{}, ctx context.Context) error
	}

	ICRUDMongoRepository[Model_T any] interface {
		Create(model *Model_T, ctx context.Context) error
		FindMany(query bson.D, ctx context.Context, projection ...bson.E) ([]*Model_T, error)
		FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*Model_T, error)
		Find(query bson.D, ctx context.Context) (*Model_T, error)
		FindOffset(
			query interface{}, offset uint64, size uint64, sort *bson.D, ctx context.Context, projection ...bson.E,
		) ([]Model_T, error)
		UpdateOneByUUID(uuid uuid.UUID, model *Model_T, ctx context.Context) error
		UpdateManyByFilter(filter bson.D, update bson.D, ctx context.Context) error
		UpsertManyByFilter(filter bson.D, update bson.D, ctx context.Context) error
		DeleteMany(model *Model_T, ctx context.Context) error
		DeleteManyByFilter(filter bson.D, ctx context.Context) error
	}

	crud_mongo_repository[Model_T any] struct {
		MongoDBQueryMonitorCollection
		//collection *mongo.Collection
	}
)

func (this *crud_mongo_repository[Model_T]) InitCollection(col *mongo.Collection) {

	this.collection = col
}

func (this *crud_mongo_repository[Model_T]) Create(model *Model_T, ctx context.Context) error {

	return createDocument[Model_T](model, &this.MongoDBQueryMonitorCollection, ctx)
}

func (this *crud_mongo_repository[Model_T]) Find(query bson.D, ctx context.Context) (*Model_T, error) {

	ret, err := findOneDocument[Model_T](query, &this.MongoDBQueryMonitorCollection, ctx)

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *crud_mongo_repository[Model_T]) FindMany(query bson.D, ctx context.Context, projections ...bson.E) ([]*Model_T, error) {

	ret, err := findManyDocuments[Model_T](query, &this.MongoDBQueryMonitorCollection, ctx, projections...)

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *crud_mongo_repository[Model_T]) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*Model_T, error) {

	ret, err := findDocumentByUUID[Model_T](uuid, &this.MongoDBQueryMonitorCollection, ctx)

	if errors.Is(err, mongo.ErrNoDocuments) {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *crud_mongo_repository[Model_T]) FindOffset(
	query interface{}, offset uint64, size uint64, sort *bson.D, ctx context.Context, projection ...bson.E,
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

func (this *crud_mongo_repository[Model_T]) UpdateOneByUUID(uuid uuid.UUID, model *Model_T, ctx context.Context) error {

	res, err := updateDocument[Model_T](libCommon.PointerPrimitive(uuid), model, &this.MongoDBQueryMonitorCollection, ctx)

	if err != nil {

		return err
	}

	return CheckUpdateOneResult(res)
}

func (this *crud_mongo_repository[Model_T]) UpdateManyByFilter(filter bson.D, update bson.D, ctx context.Context) error {

	res, err := this.MongoDBQueryMonitorCollection.UpdateMany(ctx, filter, bson.D{{"$set", update}})

	if err != nil {

		return err
	}

	return CheckUpdateOneResult(res)
}

func (this *crud_mongo_repository[Model_T]) UpsertManyByFilter(filter bson.D, update bson.D, ctx context.Context) error {

	opts := options.Update().SetUpsert(true)

	_, err := this.MongoDBQueryMonitorCollection.UpdateMany(ctx, filter, bson.D{{"$set", update}}, opts)

	if err != nil {

		return err
	}

	return nil
}

func (this *crud_mongo_repository[Model_T]) DeleteMany(model *Model_T, ctx context.Context) error {

	_, err := this.MongoDBQueryMonitorCollection.DeleteMany(ctx, model)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func (this *crud_mongo_repository[Model_T]) DeleteManyByFilter(filter bson.D, ctx context.Context) error {

	_, err := this.MongoDBQueryMonitorCollection.DeleteMany(ctx, filter)

	if err != nil {

		return libError.NewInternal(err)
	}

	return nil
}

func (this *crud_mongo_repository[Model_T]) CreateMany(models []*Model_T, ctx context.Context) error {

	_, err := insertMany(models, &this.MongoDBQueryMonitorCollection, context.TODO())

	if err != nil {

		return err
	}

	return nil
}
