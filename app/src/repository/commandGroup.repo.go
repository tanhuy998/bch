package repository

import (
	"app/src/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
		crud crud_mongo_repository[model.CommandGroup]
	}
)

func (this *CommandGroupRepository) Init(db *mongo.Database) *CommandGroupRepository {

	this.AbstractMongoRepository.Init(db, COMMAND_GROUP_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *CommandGroupRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *CommandGroupRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *CommandGroupRepository) Create(model *model.CommandGroup, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *CommandGroupRepository) Find(query bson.D, ctx context.Context) (*model.CommandGroup, error) {

	return this.crud.Find(query, ctx)
}
func (this *CommandGroupRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.CommandGroup, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *CommandGroupRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.CommandGroup, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *CommandGroupRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}
