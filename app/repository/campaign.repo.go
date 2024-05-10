package repository

import (
	"app/app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CAMPAIGN_COLLECTION_NAME = "campaigns"
)

type CampaignRepository struct {
	collection *mongo.Collection
}

func (this *CampaignRepository) Init(db *mongo.Database) {

	this.collection = db.Collection(CAMPAIGN_COLLECTION_NAME)
}

func (this *CampaignRepository) FindByUUID(uuid uuid.UUID) (*model.Campaign, error) {

	res := this.collection.FindOne(context.TODO(), bson.M{
		"uuid": uuid,
	})

	var camp *model.Campaign

	err := res.Decode(&camp)

	if err != nil {

		return nil, err
	}

	return camp, nil
}

func (this *CampaignRepository) InsertCandidates(uuid uuid.UUID) {

}
