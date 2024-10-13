package repository

import (
	"app/model"

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
		crud_mongo_repository[model.User]
	}
)

func (this *UserRepository) Init(db *mongo.Database) *UserRepository {

	this.AbstractMongoRepository.Init(db, USER_COLLECTION_NAME)

	this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *UserRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *UserRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}
