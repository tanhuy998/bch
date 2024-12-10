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

	ITenant = repositoryAPI.ICRUDMongoRepository[model.Tenant]

	TenantRepository struct {
		//AbstractMongoRepository
		mongoRepository.MongoCRUDRepository[model.Tenant]
	}
)

var (
	t repositoryAPI.ICRUDMongoRepository[model.Tenant] = new(TenantRepository)
)

func (this *TenantRepository) Init(db *mongo.Database) *TenantRepository {

	this.MongoCRUDRepository.Init(db, TENANT_COLLECTION_NAME)

	return this
}
