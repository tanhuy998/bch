package repository

import (
	"app/src/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USER_COLLECTION_NAME = "users"
)

type (
	IUser interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.User]
	}

	UserRepository struct {
		AbstractMongoRepository
		crud crud_mongo_repository[model.User]
	}
)

func (this *UserRepository) Init(db *mongo.Database) *UserRepository {

	this.AbstractMongoRepository.Init(db, USER_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *UserRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *UserRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *UserRepository) Create(model *model.User, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *UserRepository) Find(query bson.D, ctx context.Context) (*model.User, error) {

	return this.crud.Find(query, ctx)
}
func (this *UserRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.User, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *UserRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.User, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *UserRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}
