package repository

import (
	"app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	USER_SESSION_COLLECTION_NAME = "userSessions"
)

type (
	// IUserSession interface {
	// 	IMongoDBRepository
	// 	ICRUDMongoRepository[model.UserSession]
	// 	//CreateMany(models []*model.UserSession, ctx context.Context) error
	// 	ICreateMany[model.UserSession]
	// }

	IUserSession = IRepository[model.UserSession]

	UserSessionRepository struct {
		AbstractMongoRepository
		crud_mongo_repository[model.UserSession]
	}
)

func (this *UserSessionRepository) Init(db *mongo.Database) *UserSessionRepository {

	// this.AbstractMongoRepository.Init(db, USER_SESSION_COLLECTION_NAME)

	// this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	this.crud_mongo_repository.Init(db, USER_COLLECTION_NAME)

	return this
}

// func (this *UserSessionRepository) GetCollection() *mongo.Collection {

// 	return this.AbstractMongoRepository.collection
// }

// func (this *UserSessionRepository) GetDBClient() *mongo.Client {

// 	return this.GetCollection().Database().Client()
// }
