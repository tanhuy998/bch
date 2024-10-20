package repository

import (
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"app/model"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ITEM_PER_PAGE = 10
)

var (
	ERR_UPDATE_NO_MATCH       error = errors.New("no match document to update")
	NOTHING_CHANGED_ON_UPDATE error = errors.New("nothing changed")
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

	IMongoDBAggregator[Model_T any] interface {
		Aggregate(pipeline mongo.Pipeline, ctx context.Context) ([]*Model_T, error)
	}

	IMongodDBCustomPagination[Model_T any] interface {
		RetrieveCustomPagination(
			pipeline mongo.Pipeline,
			paginationPivotField string,
			pivotValue interface{},
			pageLimit int64,
			isPrevDir bool,
			ctx context.Context,
		) (*PaginationPack[Model_T], error)
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
		IMongoDBAggregator[model.Candidate]
		IMongoDBRepository
		//IMongodDBCustomPagination[model.Candidate]
		Find(query bson.D, ctx context.Context) (*model.Candidate, error)
		FindByUUID(uuid.UUID, context.Context) (*model.Candidate, error)
		Get(page int, ctx context.Context) ([]*model.Candidate, error)
		Create(*model.Candidate, context.Context) error
		GetOneSigningInfo(query bson.D, ctx context.Context, projections ...bson.E) (*model.CandidateSigningInfo, error)
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

func CheckUpdateOneResult(result *mongo.UpdateResult) error {

	if result.MatchedCount < 1 {

		return libError.NewInternal(ERR_UPDATE_NO_MATCH)

	} else if result.ModifiedCount == 0 {

		return libError.NewInternal(NOTHING_CHANGED_ON_UPDATE)
	}

	return nil
}
