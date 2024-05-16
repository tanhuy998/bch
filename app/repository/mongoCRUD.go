package repository

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func getDocuments[T any](page int64, collection *mongo.Collection) ([]*T, error) {

	cursor, err := collection.Aggregate(context.TODO(), bson.D{
		{"$skip", page},
		{"$limit", ITEM_PER_PAGE},
	})

	if err != nil {

		return nil, err
	}

	return ParseCursor[T](cursor)
}

func findDocumentByUUID[T any](uuid uuid.UUID, collection *mongo.Collection) (*T, error) {

	res := collection.FindOne(context.TODO(), bson.M{
		"uuid": uuid,
	})

	var camp *T

	err := res.Decode(&camp)

	if err != nil {

		return nil, err
	}

	return camp, nil
}

func createDocument[T any](model *T, collection *mongo.Collection) error {

	//(*model).UUID = uuid.New()

	_, err := collection.InsertOne(context.TODO(), model)

	if err != nil {

		return err
	}

	return nil
}

func updateDocument[T any](uuid uuid.UUID, model *T, collection *mongo.Collection) error {

	collection.FindOneAndUpdate(context.TODO(), bson.D{{"uuid", uuid}}, model)

	return nil
}

func deleteDocument(uuid uuid.UUID, collection *mongo.Collection) error {

	_, err := collection.DeleteOne(context.TODO(), bson.D{{"uuid", uuid}})

	if err != nil {

		return err
	}

	return nil
}
