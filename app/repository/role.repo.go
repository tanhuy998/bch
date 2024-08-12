package repository

import (
	"app/domain/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ROLE_COLLECTION_NAME = "roles"
)

type (
	IRole interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.Role]
	}

	RoleRepository struct {
		AbstractMongoRepository
		crud crud_mongo_repository[model.Role]
	}
)

func (this *RoleRepository) Init(db *mongo.Database) *RoleRepository {

	this.AbstractMongoRepository.Init(db, ROLE_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *RoleRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *RoleRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *RoleRepository) Create(model *model.Role, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *RoleRepository) Find(query bson.D, ctx context.Context) (*model.Role, error) {

	return this.crud.Find(query, ctx)
}
func (this *RoleRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.Role, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *RoleRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.Role, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *RoleRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}
