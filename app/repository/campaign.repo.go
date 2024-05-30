package repository

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CAMPAIGN_COLLECTION_NAME = "campaigns"
)

type CampaignRepository struct {
	AbstractRepository
	//collection *mongo.Collection
}

func (this *CampaignRepository) Init(db *mongo.Database) *CampaignRepository {

	this.AbstractRepository.Init(db, CAMPAIGN_COLLECTION_NAME)

	return this
}

func (this *CampaignRepository) GetCollection() *mongo.Collection {

	return this.collection
}

func (this *CampaignRepository) GetDBClient() *mongo.Client {

	return this.collection.Database().Client()
}

func (this *CampaignRepository) FindByUUID(uuid uuid.UUID, ctx context.Context) (*model.Campaign, error) {

	return findDocumentByUUID[model.Campaign](uuid, this.collection, ctx)
}

func (this *CampaignRepository) Get(page int, ctx context.Context) ([]*model.Campaign, error) {

	calcPage := this.returnPageThresholdIfOutOfRange(int64(page))

	return getDocuments[model.Campaign](calcPage, this.collection, ctx)
}

func (this *CampaignRepository) GetPendingCampaigns(
	_id primitive.ObjectID,
	pageLimit int64,
	isPrevDir bool,
	ctx context.Context,
) (data *PaginationPack[model.Campaign], err error) {

	// cursor, err := this.collection.Aggregate(context.TODO(), bson.D{
	// 	{
	// 		"$search", bson.D{
	// 			{"index", "issueTime_index"},
	// 			{"searchBefore", "$$NOW"},
	// 		},
	// 	},
	// 	{"$kip", int64(page)},
	// 	{"$limit", ITEM_PER_PAGE},
	// 	{"$sort", bson.D{{"issueTime", -1}}},
	// })

	// if err != nil {

	// 	return nil, err
	// }

	// return ParseCursor[model.Campaign](cursor, ctx)

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	ret, err := getDocumentsPageByID[model.Campaign](_id, pageLimit, isPrevDir, nil, this.collection, ctx,
		bson.E{
			"expire", bson.D{
				{OP_GT, time.Now()},
			},
		},
	)

	if err != nil {

		panic(err)
	}

	return ret, nil
}

func (this *CampaignRepository) Create(model *model.Campaign, ctx context.Context) error {

	return createDocument(model, this.collection, ctx)
}

func (this *CampaignRepository) Update(model *model.Campaign, ctx context.Context) error {

	return updateDocument(model.UUID, model, this.collection, ctx)
}

func (this *CampaignRepository) Delete(uuid uuid.UUID, ctx context.Context) error {

	return deleteDocument(uuid, this.collection, ctx)
}

func (this *CampaignRepository) GetCampaignList(
	_id primitive.ObjectID,
	pageLimit int64,
	isPrevDir bool,
	ctx context.Context,
) (data *PaginationPack[model.Campaign], err error) {

	//page = this.returnPageThresholdIfOutOfRange(page)
	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)
	//return getDocuments[model.Campaign](page, this.collection, nil)
	ret, err := getDocumentsPageByID[model.Campaign](_id, pageLimit, isPrevDir, nil, this.collection, ctx)

	if err != nil {

		panic(err)
	}

	return ret, nil
}
