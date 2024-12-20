package repository

import (
	"app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COMMAND_GROUP_COLLECTION_NAME = "commandGroups"
)

type (
	ICommandGroup interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.CommandGroup]
	}

	CommandGroupRepository struct {
		AbstractMongoRepository
		crud_mongo_repository[model.CommandGroup]
	}
)

func (this *CommandGroupRepository) Init(db *mongo.Database) *CommandGroupRepository {

	this.AbstractMongoRepository.Init(db, COMMAND_GROUP_COLLECTION_NAME)

	this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *CommandGroupRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *CommandGroupRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}
