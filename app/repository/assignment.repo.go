package repository

import (
	"app/model"
	repositoryAPI "app/repository/api"
	mongoRepository "app/repository/driver/mongod"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ASSIGNMENT_COLLECTION_NAME = "assignments"
)

type (
	// IAssignment interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.Assignment]
	// 	//CreateMany(models []*model.Assignment, ctx context.Context) error
	// 	ICreateMany[model.Assignment]
	// }

	IAssignment = repositoryAPI.ICRUDMongoRepository[model.Assignment] // IRepository[model.Assignment]

	AssignmentRepository struct {
		//AbstractMongoRepository
		//crud_mongo_repository[model.Assignment]
		mongoRepository.MongoCRUDRepository[model.Assignment]
	}
)

func (this *AssignmentRepository) Init(db *mongo.Database) *AssignmentRepository {

	// this.AbstractMongoRepository.Init(db, ASSIGNMENT_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	//this.crud_mongo_repository.Init(db, ASSIGNMENT_COLLECTION_NAME)

	this.MongoCRUDRepository.Init(db, ASSIGNMENT_COLLECTION_NAME)

	return this
}

// func (this *AssignmentRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *AssignmentRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
