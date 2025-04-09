package repository

import (
	"app/model"
	repositoryAPI "app/repository/api"
	mongoRepository "app/repository/driver/mongod"

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

	IAssignmentGroup = repositoryAPI.ICRUDRepository[model.AssignmentGroup] // mongoRepository.ICRUDMongoRepository[model.AssignmentGroup] // repositoryAPI.ICRUDMongoRepository[model.AssignmentGroup] //IRepository[model.AssignmentGroup]

	AssignmentGroupRepository struct {
		//AbstractMongoRepository
		//crud_mongo_repository[model.AssignmentGroup]
		mongoRepository.MongoCRUDRepository[model.AssignmentGroup]
	}
)

func (this *AssignmentGroupRepository) Init(db *mongo.Database) *AssignmentGroupRepository {

	// this.AbstractMongoRepository.Init(db, ASSIGNMENT_GROUP_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	this.MongoCRUDRepository.Init(db, ASSIGNMENT_GROUP_COLLECTION_NAME)

	return this
}

// func (this *AssignmentGroupRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *AssignmentGroupRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
