package repository

import (
	libCommon "app/app/lib/common"
	"app/app/model"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ITEM_PER_PAGE = 10
)

type (
	UUIDModel interface {
		model.Campaign | model.Candidate
	}

	IRepository interface {
		Init(*mongo.Database)
	}

	ICampaignRepository interface {
		FindByUUID(uuid.UUID, context.Context) (*model.Campaign, error)
		Get(page int, ctx context.Context) ([]*model.Campaign, error)
		GetPendingCampaigns(page int, ctx context.Context) ([]*model.Campaign, error)
		Create(*model.Campaign, context.Context) error
		//CreateMany([]*model.Campaign) error
		Update(*model.Campaign, context.Context) error
		Delete(uuid.UUID, context.Context) error
		//Remove(uuid uuid.UUID) (bool, error)
	}

	ICandidateRepository interface {
		FindByUUID(uuid.UUID, context.Context) (*model.Campaign, error)
		Get(page int, ctx context.Context) ([]*model.Campaign, error)
		Create(*model.Candidate, context.Context) error
		Update(*model.Candidate, context.Context) error
		Delete(uuid.UUID, context.Context) error
		//Remove(uuid uuid.UUID) (bool, error)
	}

	AbstractRepository struct {
		collection *mongo.Collection
	}
)

func (this *AbstractRepository) Init(db *mongo.Database, collectionName string) {

	this.collection = db.Collection(collectionName)
}

func (this *AbstractRepository) CountPage() (int64, error) {

	docNum, err := this.collection.CountDocuments(context.TODO(), struct{}{})

	if err != nil {

		return -1, err
	}

	even := docNum / int64(ITEM_PER_PAGE)
	odd := libCommon.Ternary[int64](docNum%ITEM_PER_PAGE > 0, int64(64), int64(0))

	return even + odd, nil
}

func (this *AbstractRepository) returnPageThresholdIfOutOfRange(inputPageNum int64) int64 {

	inputPageNum = libCommon.Ternary(inputPageNum <= 0, 1, inputPageNum)

	pageCount, err := this.CountPage()

	if err != nil {

		return 1
	}

	return libCommon.Ternary[int64](inputPageNum > pageCount, pageCount, inputPageNum)
}

func (this *AbstractRepository) Collection() *mongo.Collection {

	return this.collection
}
