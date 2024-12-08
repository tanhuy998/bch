package repository

import (
	"app/model"
	repositoryAPI "app/repository/api"
	mongoRepository "app/repository/driver/mongod"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TENANT_COLLECTION_NAME = "tenants"
)

type (
	// ITenant interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.Tenant]
	// }

	ITenant = repositoryAPI.ICRUDRepository[model.Tenant]

	TenantRepository struct {
		//AbstractMongoRepository
		mongoRepository.MongoCRUDRepository[model.Tenant]
	}
)

func (this *TenantRepository) Init(db *mongo.Database) *TenantRepository {

	// this.AbstractMongoRepository.Init(db, TENANT_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	this.MongoCRUDRepository.Init(db, TENANT_COLLECTION_NAME)

	return this
}
