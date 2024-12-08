package mongoRepository

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
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ITEM_PER_PAGE = 10
)

var (
	ERR_UPDATE_NO_MATCH       error = errors.New("no match document to update")
	NOTHING_CHANGED_ON_UPDATE error = errors.New("nothing changed")
)

type (
	IMongoRepositoryOperator interface {
		Database() *mongo.Database
		BulkWrite(ctx context.Context, models []mongo.WriteModel,
			opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
		InsertOne(ctx context.Context, document interface{},
			opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
		InsertMany(ctx context.Context, documents []interface{},
			opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
		DeleteOne(ctx context.Context, filter interface{},
			opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
		DeleteMany(ctx context.Context, filter interface{},
			opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
		UpdateByID(ctx context.Context, id interface{}, update interface{},
			opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		UpdateOne(ctx context.Context, filter interface{}, update interface{},
			opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		UpdateMany(ctx context.Context, filter interface{}, update interface{},
			opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
		ReplaceOne(ctx context.Context, filter interface{},
			replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
		Aggregate(ctx context.Context, pipeline interface{},
			opts ...*options.AggregateOptions) (*mongo.Cursor, error)
		CountDocuments(ctx context.Context, filter interface{},
			opts ...*options.CountOptions) (int64, error)
		EstimatedDocumentCount(ctx context.Context,
			opts ...*options.EstimatedDocumentCountOptions) (int64, error)
		Distinct(ctx context.Context, fieldName string, filter interface{},
			opts ...*options.DistinctOptions) ([]interface{}, error)
		Find(ctx context.Context, filter interface{},
			opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
		FindOne(ctx context.Context, filter interface{},
			opts ...*options.FindOneOptions) *mongo.SingleResult
		FindOneAndDelete(ctx context.Context, filter interface{},
			opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
		FindOneAndReplace(ctx context.Context, filter interface{},
			replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult
		FindOneAndUpdate(ctx context.Context, filter interface{},
			update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
		Watch(ctx context.Context, pipeline interface{},
			opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
	}
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
		GetCollection() IMongoRepositoryOperator
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

type (
	IRepository[Model_T any] interface {
		IMongoDBRepository
		ICRUDMongoRepository[Model_T]
	}
)

type (
	ITransactionDBClient interface {
		WithTransaction(func(ctx context.Context) (interface{}, error)) (interface{}, error)
	}

	MongoDBClient struct {
		dbClient *mongo.Client
	}
)

func (this *MongoDBClient) WithTransaction(ctx context.Context, fn func(sessionCtx context.Context) (interface{}, error)) (interface{}, error) {

	session, err := this.dbClient.StartSession()

	if err != nil {

		return nil, err
	}

	defer session.EndSession(context.TODO())

	return session.WithTransaction(ctx, func(ssCtx mongo.SessionContext) (interface{}, error) {

		return fn(ssCtx)
	})
}

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

func (this *AbstractMongoRepository) GetDBClient() *mongo.Client /**tv(DBClient_T)*/ {

	return this.collection.Database().Client()
}

func CheckUpdateOneResult(result *mongo.UpdateResult) error {

	if result.MatchedCount < 1 {

		return libError.NewInternal(ERR_UPDATE_NO_MATCH)

	} else if result.ModifiedCount == 0 {

		return libError.NewInternal(NOTHING_CHANGED_ON_UPDATE)
	}

	return nil
}
