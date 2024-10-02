package repository

import (
	"app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TENANT_COLLECTION_NAME = "tenants"
)

type (
	ITenant interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.Tenant]
	}

	TenantRepository struct {
		AbstractMongoRepository
		crud crud_mongo_repository[model.Tenant]
	}
)

func (this *TenantRepository) Init(db *mongo.Database) *TenantRepository {

	this.AbstractMongoRepository.Init(db, TENANT_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *TenantRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *TenantRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *TenantRepository) Create(model *model.Tenant, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *TenantRepository) Find(query bson.D, ctx context.Context) (*model.Tenant, error) {

	return this.crud.Find(query, ctx)
}
func (this *TenantRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.Tenant, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *TenantRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.Tenant, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *TenantRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}
