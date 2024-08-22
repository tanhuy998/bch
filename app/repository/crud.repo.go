package repository

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IFindMany[Model_T any] interface {
		FindMany(query bson.D, ctx context.Context, projection ...bson.E) ([]*Model_T, error)
	}

	ICreateMany[Model_T any] interface {
		CreateMany(models []*Model_T, ctx context.Context) error
	}

	ICRUDMongoRepository[Model_T any] interface {
		Create(model *Model_T, ctx context.Context) error
		Find(query bson.D, ctx context.Context) (*Model_T, error)
		FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*Model_T, error)
		UpdateOneByUUID(uuid uuid.UUID, model *Model_T, ctx context.Context) error
		Delete(query bson.D, ctx context.Context) error
	}

	crud_mongo_repository[Model_T any] struct {
		collection *mongo.Collection
	}
)

func (this *crud_mongo_repository[Model_T]) InitCollection(col *mongo.Collection) {

	this.collection = col
}

func (this *crud_mongo_repository[Model_T]) Create(model *Model_T, ctx context.Context) error {

	return createDocument[Model_T](model, this.collection, ctx)
}

func (this *crud_mongo_repository[Model_T]) Find(query bson.D, ctx context.Context) (*Model_T, error) {

	ret, err := findOneDocument[Model_T](query, this.collection, ctx)

	if err == mongo.ErrNoDocuments {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *crud_mongo_repository[Model_T]) FindMany(query bson.D, ctx context.Context, projections ...bson.E) ([]*Model_T, error) {

	ret, err := findManyDocuments[Model_T](query, this.collection, ctx, projections...)

	if err == mongo.ErrNoDocuments {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *crud_mongo_repository[Model_T]) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*Model_T, error) {

	ret, err := findDocumentByUUID[Model_T](uuid, this.collection, ctx)

	if err == mongo.ErrNoDocuments {

		return nil, nil
	}

	if err != nil {

		return nil, err
	}

	return ret, nil
}

// func (this *crud_mongo_repository[Model_T]) FindByUUID(uuid uuid.UUID, ctx context.Context) ([]*Model_T, error) {

// 	return nil, nil
// }

func (this *crud_mongo_repository[Model_T]) UpdateOneByUUID(uuid uuid.UUID, model *Model_T, ctx context.Context) error {

	res, err := updateDocument[Model_T](libCommon.PointerPrimitive(uuid), model, this.collection, ctx)

	if err != nil {

		return err
	}

	return CheckUpdateOneResult(res)
}

func (this *crud_mongo_repository[Model_T]) Delete(query bson.D, ctx context.Context) error {

	return nil
}

func (this *crud_mongo_repository[Model_T]) CreateMany(models []*model.Role, ctx context.Context) error {

	_, err := insertMany(models, this.collection, context.TODO())

	if err != nil {

		return err
	}

	return nil
}
