package repository

import (
	"app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COMMAND_GROUP_USER_ROLE_COLLECTION_NAME = "commandGroupUserRoles"
)

type (
	ICommandGroupUserRole interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.CommandGroupUserRole]
		//CreateMany(models []*model.CommandGroupUserRole, ctx context.Context) error
		ICreateMany[model.CommandGroupUserRole]
	}

	CommandGroupUserRoleRepository struct {
		AbstractMongoRepository
		crud_mongo_repository[model.CommandGroupUserRole]
	}
)

func (this *CommandGroupUserRoleRepository) Init(db *mongo.Database) *CommandGroupUserRoleRepository {

	this.AbstractMongoRepository.Init(db, COMMAND_GROUP_USER_ROLE_COLLECTION_NAME)

	this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *CommandGroupUserRoleRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *CommandGroupUserRoleRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}
