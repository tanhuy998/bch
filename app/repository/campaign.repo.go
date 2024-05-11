package repository

import (
	"app/app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	ITEM_PER_PAGE            = 10
	CAMPAIGN_COLLECTION_NAME = "campaigns"
)

type CampaignRepository struct {
	IRepository
	AbstractRepository
	//collection *mongo.Collection
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

func (this *CandidateRepository) RetrieveCandidatesOfCampaign(campaign uuid.UUID, page int) ([]model.Candidate, error) {

	if page <= 0 {

		page = 1
	}

	coll := this.collection

	//coll.CountDocuments()
	convertedPageNum := this.returnPageThresholdIfOutOfRange(int64(page))

	cursor, err := coll.Aggregate(
		context.TODO(),
		bson.D{
			// {
			// 	"$match", bson.D{
			// 		{"uuid", campaign},
			// 	},
			// },
			// {"$unwind", "candidates"},
			// {"$skip", convertedPageNum * int64(ITEM_PER_PAGE)},
			// {"$limit", ITEM_PER_PAGE},
			// {
			// 	"$lookup", bson.D{
			// 		{"from", "candidates"},
			// 		{"localField", "candidate_ids"},
			// 		{"foreignField", "_id"},
			// 		{"as", "detail"},
			// 	},
			// },
			// {
			// 	"project", bson.D{
			// 		{"$detail.name", 1},
			// 		{"$detail.address", 1},
			// 		{"$detail.idNumber", 1},
			// 	},
			// },
		},
	)

	if err != nil {

		return nil, err
	}

	parsedList, err := ParseCursor[struct {
		Detail model.Candidate `bson:"detail"`
	}](cursor)

	if err != nil {

		return nil, err
	}

	var ret []model.Candidate = []model.Candidate{}

	for _, model := range parsedList {

		ret = append(ret, model.Detail)
	}

	return ret, nil
}

func (this *CampaignRepository) InsertCandidates(uuid uuid.UUID) {

}
