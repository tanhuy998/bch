package repository

import (
	"app/model"
	repositoryAPI "app/repository/api"
	mongoRepository "app/repository/driver/mongod"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COMMAND_GROUP_USER_ROLE_COLLECTION_NAME = "commandGroupUserRoles"
)

type (
	// ICommandGroupUserRole interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.CommandGroupUserRole]
	// 	//CreateMany(models []*model.CommandGroupUserRole, ctx context.Context) error
	// 	ICreateMany[model.CommandGroupUserRole]
	// }

	ICommandGroupUserRole = repositoryAPI.ICRUDRepository[model.CommandGroupUserRole] //IRepository[model.CommandGroupUserRole]

	CommandGroupUserRoleRepository struct {
		//AbstractMongoRepository
		//crud_mongo_repository[model.CommandGroupUserRole]
		mongoRepository.MongoCRUDRepository[model.CommandGroupUserRole]
	}
)

func (this *CommandGroupUserRoleRepository) Init(db *mongo.Database) *CommandGroupUserRoleRepository {

	// this.AbstractMongoRepository.Init(db, COMMAND_GROUP_USER_ROLE_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	//this.crud_mongo_repository.Init(db, COMMAND_GROUP_USER_ROLE_COLLECTION_NAME)

	this.MongoCRUDRepository.Init(db, COMMAND_GROUP_USER_ROLE_COLLECTION_NAME)

	return this
}

// func (this *CommandGroupUserRoleRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *CommandGroupUserRoleRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
