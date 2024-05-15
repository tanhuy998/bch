package repository

import (
	"app/app/model"

	"github.com/google/uuid"
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

func (this *CampaignRepository) FindByUUID(uuid uuid.UUID) (*model.Campaign, error) {

	return findDocumentByUUID[model.Campaign](uuid, this.collection)
}

func (this *CampaignRepository) Get(page int) ([]*model.Campaign, error) {

	calcPage := this.returnPageThresholdIfOutOfRange(int64(page))

	return getDocuments[model.Campaign](calcPage, this.collection)
}

func (this *CampaignRepository) GetPendingCampaigns(page int) ([]*model.Campaign, error) {

	return make([]*model.Campaign, 0), nil
}

func (this *CampaignRepository) Create(model *model.Campaign) error {

	model.UUID = uuid.New()

	return createDocument(model, this.collection)
}

func (this *CampaignRepository) Update(model *model.Campaign) error {

	return updateDocument(model.UUID, model, this.collection)
}

func (this *CampaignRepository) Delete(uuid uuid.UUID) error {

	return deleteDocument(uuid, this.collection)
}
