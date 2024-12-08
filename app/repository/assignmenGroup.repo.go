package repository

import (
	"app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ASSIGNMENT_GROUP_COLLECTION_NAME = "assignmentGroups"
)

type (
	// IAssignmentGroup interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.AssignmentGroup]
	// 	//CreateMany(models []*model.AssignmentGroup, ctx context.Context) error
	// 	ICreateMany[model.AssignmentGroup]
	// }

	IAssignmentGroup = IRepository[model.AssignmentGroup]

	AssignmentGroupRepository struct {
		//AbstractMongoRepository
		crud_mongo_repository[model.AssignmentGroup]
	}
)

func (this *AssignmentGroupRepository) Init(db *mongo.Database) *AssignmentGroupRepository {

	// this.AbstractMongoRepository.Init(db, ASSIGNMENT_GROUP_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	this.crud_mongo_repository.Init(db, ASSIGNMENT_GROUP_COLLECTION_NAME)

	return this
}

// func (this *AssignmentGroupRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *AssignmentGroupRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
