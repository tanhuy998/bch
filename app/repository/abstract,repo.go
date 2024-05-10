package repository

import (
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

func ParseCursor[T any](cursor *mongo.Cursor, val T) ([]*T, error) {

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
