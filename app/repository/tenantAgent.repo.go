package repository

import (
	"app/model"
	repositoryAPI "app/repository/api"
	mongoRepository "app/repository/driver/mongod"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TENANT_AGENT_COLLECTION_NAME = "tenantAgents"
)

type (
	// ITenantAgent interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.TenantAgent]
	// }

	ITenantAgent = repositoryAPI.ICRUDRepository[model.TenantAgent] // IRepository[model.TenantAgent]

	TenantAgentRepository struct {
		//AbstractMongoRepository
		//crud_mongo_repository[model.TenantAgent]
		mongoRepository.MongoCRUDRepository[model.TenantAgent]
	}
)

func (this *TenantAgentRepository) Init(db *mongo.Database) *TenantAgentRepository {

	// this.AbstractMongoRepository.Init(db, TENANT_AGENT_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	// this.crud_mongo_repository.Init(db, TENANT_AGENT_COLLECTION_NAME)

	this.MongoCRUDRepository.Init(db, TENANT_AGENT_COLLECTION_NAME)

	return this
}

// func (this *TenantAgentRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *TenantAgentRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
