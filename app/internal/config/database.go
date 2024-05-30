package config

import (
	libCommon "app/lib/common"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDomainIndexes(db *mongo.Database) {

	db.Collection(CAMPAIGN_COLLETCTION).Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			mongo.IndexModel{
				Keys: "uuid",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("uuid"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
		},
	)
	db.Collection(CANDIATE_COLLECTIONN).Indexes().CreateMany(
		context.TODO(),
		[]mongo.IndexModel{
			mongo.IndexModel{
				Keys: "uuid",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("uuid"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
			mongo.IndexModel{
				Keys: "idNumber",
				Options: &options.IndexOptions{
					Name:   libCommon.PointerPrimitive("idNumber"),
					Unique: libCommon.PointerPrimitive(true),
				},
			},
		},
	)
}
