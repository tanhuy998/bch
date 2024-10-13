package repository

import (
	"app/model"

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
		crud_mongo_repository[model.Tenant]
	}
)

func (this *TenantRepository) Init(db *mongo.Database) *TenantRepository {

	this.AbstractMongoRepository.Init(db, TENANT_COLLECTION_NAME)

	this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *TenantRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *TenantRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}
