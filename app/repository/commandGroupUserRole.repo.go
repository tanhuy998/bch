package repository

import (
	"app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
		crud crud_mongo_repository[model.CommandGroupUserRole]
	}
)

func (this *CommandGroupUserRoleRepository) Init(db *mongo.Database) *CommandGroupUserRoleRepository {

	this.AbstractMongoRepository.Init(db, COMMAND_GROUP_USER_ROLE_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *CommandGroupUserRoleRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *CommandGroupUserRoleRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *CommandGroupUserRoleRepository) Create(model *model.CommandGroupUserRole, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *CommandGroupUserRoleRepository) Find(query bson.D, ctx context.Context) (*model.CommandGroupUserRole, error) {

	return this.crud.Find(query, ctx)
}
func (this *CommandGroupUserRoleRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.CommandGroupUserRole, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *CommandGroupUserRoleRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.CommandGroupUserRole, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *CommandGroupUserRoleRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}

func (this *CommandGroupUserRoleRepository) CreateMany(
	models []*model.CommandGroupUserRole,
	ctx context.Context,
) error {

	_, err := insertMany(models, this.collection, context.TODO())

	if err != nil {

		return err
	}

	return nil
}
