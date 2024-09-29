package repository

import (
	"app/domain/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COMMAND_GROUP_USER_COLLECTION_NAME = "commandGroupUsers"
)

type (
	ICommandGroupUser interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.CommandGroupUser]
	}

	CommandGroupUserRepository struct {
		AbstractMongoRepository
		crud crud_mongo_repository[model.CommandGroupUser]
	}
)

func (this *CommandGroupUserRepository) Init(db *mongo.Database) *CommandGroupUserRepository {

	this.AbstractMongoRepository.Init(db, COMMAND_GROUP_USER_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *CommandGroupUserRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *CommandGroupUserRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *CommandGroupUserRepository) Create(model *model.CommandGroupUser, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *CommandGroupUserRepository) Find(query bson.D, ctx context.Context) (*model.CommandGroupUser, error) {

	return this.crud.Find(query, ctx)
}
func (this *CommandGroupUserRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.CommandGroupUser, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *CommandGroupUserRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.CommandGroupUser, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *CommandGroupUserRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}
