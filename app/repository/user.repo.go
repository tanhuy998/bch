package repository

import (
	"app/model"
	repositoryAPI "app/repository/api"
	mongoRepository "app/repository/driver/mongod"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USER_COLLECTION_NAME = "users"
)

type (
	// IUser interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.User]
	// }

	// IUser = IRepository[model.User]

	IUser = repositoryAPI.ICRUDRepository[model.User]

	UserRepository struct {
		//AbstractMongoRepository
		// crud_mongo_repository[model.User]
		mongoRepository.MongoCRUDRepository[model.User]
	}
)

func (this *UserRepository) Init(db *mongo.Database) *UserRepository {

	// this.AbstractMongoRepository.Init(db, USER_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	// this.crud_mongo_repository.Init(db, USER_COLLECTION_NAME)

	this.MongoCRUDRepository.Init(db, USER_COLLECTION_NAME)

	return this
}

// func (this *UserRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *UserRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
