package repository

import (
	"app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ASSIGNMENT_GROUP_MEMBER_COLLECTION_NAME = "assignmentGroupMembers"
)

type (
	IAssignmentGroupMember interface {
		IMongoDBRepository
		ICRUDMongoRepository[model.AssignmentGroupMember]
		//CreateMany(models []*model.AssignmentGroupMember, ctx context.Context) error
		ICreateMany[model.AssignmentGroupMember]
	}

	AssignmentGroupMemberRepository struct {
		AbstractMongoRepository
		crud_mongo_repository[model.AssignmentGroupMember]
	}
)

func (this *AssignmentGroupMemberRepository) Init(db *mongo.Database) *AssignmentGroupMemberRepository {

	this.AbstractMongoRepository.Init(db, ASSIGNMENT_GROUP_MEMBER_COLLECTION_NAME)

	this.crud_mongo_repository.InitCollection(this.AbstractMongoRepository.collection)

	return this
}

func (this *AssignmentGroupMemberRepository) GetCollection() *mongo.Collection {

	return this.AbstractMongoRepository.collection
}

func (this *AssignmentGroupMemberRepository) GetDBClient() *mongo.Client {

	return this.GetCollection().Database().Client()
}

// func (this *AssignmentGroupMemberRepository) Create(model *model.AssignmentGroupMember, ctx context.Context) error {

// 	return this.crud.Create(model, ctx)
// }

// func (this *AssignmentGroupMemberRepository) Find(query bson.D, ctx context.Context) (*model.AssignmentGroupMember, error) {

// 	return this.crud.Find(query, ctx)
// }
// func (this *AssignmentGroupMemberRepository) FindOneByUUID(uuid uuid.UUID, ctx context.Context) (*model.AssignmentGroupMember, error) {

// 	return this.crud.FindOneByUUID(uuid, ctx)
// }

// func (this *AssignmentGroupMemberRepository) UpdateOneByUUID(uuid uuid.UUID, model *model.AssignmentGroupMember, ctx context.Context) error {

// 	return this.crud.UpdateOneByUUID(uuid, model, ctx)
// }

// func (this *AssignmentGroupMemberRepository) Delete(query bson.D, ctx context.Context) error {

// 	return this.crud.Delete(query, ctx)
// }

// func (this *AssignmentGroupMemberRepository) CreateMany(
// 	models []*model.AssignmentGroupMember,
// 	ctx context.Context,
// ) error {

// 	_, err := insertMany(models, this.collection, context.TODO())

// 	if err != nil {

// 		return err
// 	}

// 	return nil
// }
