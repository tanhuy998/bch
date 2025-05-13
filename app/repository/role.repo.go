package repository

import (
	"app/model"
	repositoryAPI "app/repository/api"
	mongoRepository "app/repository/driver/mongod"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ROLE_COLLECTION_NAME = "roles"
)

type (
	// IRole interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.Role]
	// 	//FindMany(query bson.D, ctx context.Context) ([]*model.Role, error)
	// 	//IFindMany[model.Role]
	// 	ICreateMany[model.Role]
	// }

	IRole = repositoryAPI.ICRUDRepository[model.Role] // IRepository[model.Role]

	RoleRepository struct {
		//AbstractMongoRepository
		//crud_mongo_repository[model.Role]

		mongoRepository.MongoCRUDRepository[model.Role]
	}
)

func (this *RoleRepository) Init(db *mongo.Database) *RoleRepository {

	// this.AbstractMongoRepository.Init(db, ROLE_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	//this.crud_mongo_repository.Init(db, ROLE_COLLECTION_NAME)

	this.MongoCRUDRepository.Init(db, ROLE_COLLECTION_NAME)

	return this
}

// func (this *RoleRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *RoleRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
