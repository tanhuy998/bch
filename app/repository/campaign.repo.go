package repository

import (
	"app/domain/model"
	"context"

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

func (this *CampaignRepository) FindByUUID(uuid uuid.UUID, ctx context.Context) (*model.Campaign, error) {

	return findDocumentByUUID[model.Campaign](uuid, this.collection, ctx)
}

func (this *CampaignRepository) Get(page int, ctx context.Context) ([]*model.Campaign, error) {

	calcPage := this.returnPageThresholdIfOutOfRange(int64(page))

	return getDocuments[model.Campaign](calcPage, this.collection, ctx)
}

func (this *CampaignRepository) GetPendingCampaigns(page int, ctx context.Context) ([]*model.Campaign, error) {

	cursor, err := this.collection.Aggregate(context.TODO(), bson.D{
		{
			"$search", bson.D{
				{"index", "issueTime_index"},
				{"searchBefore", "$$NOW"},
			},
		},
		{"$kip", int64(page)},
		{"$limit", ITEM_PER_PAGE},
		{"$sort", bson.D{{"issueTime", -1}}},
	})

	if err != nil {

		return nil, err
	}

	return ParseCursor[model.Campaign](cursor, ctx)
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
	direction int64,
) ([]*model.Campaign, error) {

	//page = this.returnPageThresholdIfOutOfRange(page)

	//return getDocuments[model.Campaign](page, this.collection, nil)
	return getDocumentsPageByID[model.Campaign](_id, pageLimit, direction, nil, this.collection, nil)
}
