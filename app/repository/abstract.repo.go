package repository

import (
	lib "app/app/lib/concrete/error"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	Init(db *mongo.Database)
}

type AbstractRepository struct {
	collection *mongo.Collection
}

func (this *AbstractRepository) Init(db *mongo.Database) {

	this.collection = db.Collection(CAMPAIGN_COLLECTION_NAME)
}

func (this *AbstractRepository) CountPage() (int64, error) {

	docNum, err := this.collection.CountDocuments(context.TODO(), struct{}{})

	if err != nil {

		return -1, err
	}

	even := docNum / int64(ITEM_PER_PAGE)
	odd := lib.Ternary[int64](docNum%ITEM_PER_PAGE > 0, int64(64), int64(0))

	return even + odd, nil
}

func (this *AbstractRepository) returnPageThresholdIfOutOfRange(inputPageNum int64) int64 {

	inputPageNum = lib.Ternary(inputPageNum <= 0, 1, inputPageNum)

	pageCount, err := this.CountPage()

	if err != nil {

		return 1
	}

	return lib.Ternary[int64](inputPageNum > pageCount, pageCount, inputPageNum)
}

func ParseCursor[T any](cursor *mongo.Cursor) ([]*T, error) {

	var ret []*T

	for cursor.Next(context.TODO()) {

		var parsedModel = new(T)

		if err := cursor.Decode(&parsedModel); err != nil {

			return nil, err
		}

		ret = append(ret, parsedModel)
	}

	if err := cursor.Err(); err != nil {

		return nil, err
	}

	return ret, nil
}
