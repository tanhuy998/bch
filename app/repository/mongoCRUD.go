package repository

import (
	libCommon "app/lib/common"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseCursor[T any](cursor *mongo.Cursor, ctx context.Context) ([]*T, error) {

	var ret []*T

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	for cursor.Next(ctx) {

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

func getDocuments[T any](page int64, collection *mongo.Collection, ctx context.Context) ([]*T, error) {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	cursor, err := collection.Aggregate(ctx, bson.D{
		{"$skip", page},
		{"$limit", ITEM_PER_PAGE},
	})

	if err != nil {

		return nil, err
	}

	return ParseCursor[T](cursor, ctx)
}

func findDocumentByUUID[T any](uuid uuid.UUID, collection *mongo.Collection, ctx context.Context) (*T, error) {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	res := collection.FindOne(ctx, bson.M{
		"uuid": uuid,
	})

	var camp *T

	err := res.Decode(&camp)

	if err != nil {

		return nil, err
	}

	return camp, nil
}

func createDocument[T any](model *T, collection *mongo.Collection, ctx context.Context) error {

	//(*model).UUID = uuid.New()

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	_, err := collection.InsertOne(ctx, model)

	if err != nil {

		return err
	}

	return nil
}

func updateDocument[T any](uuid uuid.UUID, model *T, collection *mongo.Collection, ctx context.Context) error {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	result, err := collection.UpdateOne(ctx, bson.D{{"uuid", uuid}}, bson.D{{"$set", model}})

	if err != nil {

		return err
	}

	if result.ModifiedCount == 0 {

		return errors.New("No matched document to update")
	}

	return nil
}

func deleteDocument(uuid uuid.UUID, collection *mongo.Collection, ctx context.Context) error {

	ctx = libCommon.Ternary(ctx == nil, context.TODO(), ctx)

	_, err := collection.DeleteOne(ctx, bson.D{{"uuid", uuid}})

	if err != nil {

		return err
	}

	return nil
}
