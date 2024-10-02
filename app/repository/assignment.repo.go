package repository

import (
	"app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ASSIGNMENT_COLLECTION_NAME = "assignments"
)

type (
	IAssignment interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.Assignment]
		//CreateMany(models []*model.Assignment, ctx context.Context) error
		ICreateMany[model.Assignment]
	}

	AssignmentRepository struct {
		AbstractMongoRepository
		crud crud_mongo_repository[model.Assignment]
	}
)

func (this *AssignmentRepository) Init(db *mongo.Database) *AssignmentRepository {

	this.AbstractMongoRepository.Init(db, ASSIGNMENT_COLLECTION_NAME)

	this.crud.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *AssignmentRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *AssignmentRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

func (this *AssignmentRepository) Create(model *model.Assignment, ctx context.Context) error {

	return this.crud.Create(model, ctx)
}

func (this *AssignmentRepository) Find(query bson.D, ctx context.Context) (*model.Assignment, error) {

	return this.crud.Find(query, ctx)
}
func (this *AssignmentRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.Assignment, error) {

	return this.crud.FindOneByUUID(uuid, ctx)
}

func (this *AssignmentRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.Assignment, ctx context.Context) error {

	return this.crud.UpdateOneByUUID(uuid, model, ctx)
}

func (this *AssignmentRepository) Delete(query bson.D, ctx context.Context) error {

	return this.crud.Delete(query, ctx)
}

func (this *AssignmentRepository) CreateMany(
	models []*model.Assignment,
	ctx context.Context,
) error {

	_, err := insertMany(models, this.collection, context.TODO())

	if err != nil {

		return err
	}

	return nil
}
