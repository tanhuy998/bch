package repository

import (
	"app/src/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TENANT_AGENT_COLLECTION_NAME = "tenantAgents"
)

type (
	ITenantAgent interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.TenantAgent]
	}

	TenantAgentRepository struct {
		AbstractMongoRepository
		crud crud_mongo_repository[model.TenantAgent]
	}
)

func (this *TenantAgentRepository) Init(db *mongo.Database) *TenantAgentRepository {

	this.AbstractMongoRepository.Init(db, TENANT_AGENT_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *TenantAgentRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *TenantAgentRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *TenantAgentRepository) Create(model *model.TenantAgent, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *TenantAgentRepository) Find(query bson.D, ctx context.Context) (*model.TenantAgent, error) {

	return this.crud.Find(query, ctx)
}
func (this *TenantAgentRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.TenantAgent, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *TenantAgentRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.TenantAgent, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *TenantAgentRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}
