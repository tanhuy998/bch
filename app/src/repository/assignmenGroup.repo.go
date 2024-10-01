package repository

import (
	"app/src/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ASSIGNMENT_GROUP_COLLECTION_NAME = "assignmentGroups"
)

type (
	IAssignmentGroup interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.AssignmentGroup]
		//CreateMany(models []*model.AssignmentGroup, ctx context.Context) error
		ICreateMany[model.AssignmentGroup]
	}

	AssignmentGroupRepository struct {
		AbstractMongoRepository
		crud crud_mongo_repository[model.AssignmentGroup]
	}
)

func (this *AssignmentGroupRepository) Init(db *mongo.Database) *AssignmentGroupRepository {

	this.AbstractMongoRepository.Init(db, ASSIGNMENT_GROUP_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *AssignmentGroupRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *AssignmentGroupRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *AssignmentGroupRepository) Create(model *model.AssignmentGroup, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *AssignmentGroupRepository) Find(query bson.D, ctx context.Context) (*model.AssignmentGroup, error) {

	return this.crud.Find(query, ctx)
}
func (this *AssignmentGroupRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.AssignmentGroup, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *AssignmentGroupRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.AssignmentGroup, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *AssignmentGroupRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}

func (this *AssignmentGroupRepository) CreateMany(
	models []*model.AssignmentGroup,
	ctx context.Context,
) error {

	_, err := insertMany(models, this.collection, context.TODO())

	if err != nil {

		return err
	}

	return nil
}
