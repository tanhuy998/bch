package repository

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ITEM_PER_PAGE = 10
)

type (
	UUIDModel interface {
		model.Campaign | model.Candidate
	}

	IAbstractRepository[DBClient_T any] interface {
		GetDBClient() *DBClient_T
	}

	IMongoDBRepository interface {
		IAbstractRepository[mongo.Client]
		//Init(*mongo.Database)
		GetCollection() *mongo.Collection
	}

	PaginationPack[Model_T any] struct {
		Data  []*Model_T
		Count int64
	}

	PaginationPackWithHeader[Model_T, Header_T any] struct {
		Header *Header_T
		*PaginationPack[Model_T]
	}

	CandidateSigninInfoUpdateQuery struct {
		SigningInfo *model.CandidateSigningInfo `bson:"signingInfo,omitEmpty"`
	}

	ICampaignRepository interface {
		IMongoDBRepository
		FindByUUID(uuid.UUID, context.Context) (*model.Campaign, error)
		Get(page int, ctx context.Context) ([]*model.Campaign, error)
		GetPendingCampaigns(
			id primitive.ObjectID,
			pageLimit int64,
			direction bool,
			ctx context.Context,
		) (data *PaginationPack[model.Campaign], err error)
		GetCampaignList(
			id primitive.ObjectID,
			pageLimit int64,
			direction bool,
			ctx context.Context,
		) (data *PaginationPack[model.Campaign], err error)
		Create(*model.Campaign, context.Context) error
		//CreateMany([]*model.Campaign) error
		Update(*model.Campaign, context.Context) error
		Delete(uuid.UUID, context.Context) error
		//Remove(uuid uuid.UUID) (bool, error)
	}

	ICandidateRepository interface {
		IMongoDBRepository
		FindByUUID(uuid.UUID, context.Context) (*model.Candidate, error)
		Get(page int, ctx context.Context) ([]*model.Candidate, error)
		Create(*model.Candidate, context.Context) error
		Update(*model.Candidate, context.Context) error
		UpdateSigningInfo(candidateUUID uuid.UUID, campaignUUID uuid.UUID, query *CandidateSigninInfoUpdateQuery, ctx context.Context) error
		Delete(uuid.UUID, context.Context) error
		GetCandidaiteList(
			campaignUUID uuid.UUID,
			pivot_id primitive.ObjectID,
			pageLimit int64,
			isPrevDir bool,
			ctx context.Context,
		) (*PaginationPack[model.Candidate], error)
		//Remove(uuid uuid.UUID) (bool, error)
	}

	AbstractMongoRepository struct {
		collection *mongo.Collection
	}
)

func (this *AbstractMongoRepository) Init(db *mongo.Database, collectionName string) {

	this.collection = db.Collection(collectionName)
}

func (this *AbstractMongoRepository) CountPage() (int64, error) {

	docNum, err := this.collection.CountDocuments(context.TODO(), struct{}{})

	if err != nil {

		return -1, err
	}

	even := docNum / int64(ITEM_PER_PAGE)
	odd := libCommon.Ternary[int64](docNum%ITEM_PER_PAGE > 0, int64(64), int64(0))

	return even + odd, nil
}

func (this *AbstractMongoRepository) returnPageThresholdIfOutOfRange(inputPageNum int64) int64 {

	inputPageNum = libCommon.Ternary(inputPageNum <= 0, 1, inputPageNum)

	pageCount, err := this.CountPage()

	if err != nil {

		return 1
	}

	return libCommon.Ternary[int64](inputPageNum > pageCount, pageCount, inputPageNum)
}

func (this *AbstractMongoRepository) Collection() *mongo.Collection {

	return this.collection
}
