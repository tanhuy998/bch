package config

import (
	libCommon "app/lib/common"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

	err := db.CreateView(
		context.TODO(),
		"signingDetails",
		CANDIATE_COLLECTIONN,
		bson.A{
			bson.D{
				{
					"$lookup", bson.D{
						{"from", CAMPAIGN_COLLETCTION},
						{"localField", "campaignUUID"},
						{"foreignField", "uuid"},
						{"as", "campaignUUID"},
					},
				},
			},
		},
	)

	// err = db.CreateView(
	// 	context.TODO(),
	// 	"signingDetails",
	// 	CAMPAIGN_COLLETCTION,
	// 	bson.A{
	// 		bson.D{
	// 			{
	// 				"$lookup", bson.D{
	// 					{"from", CANDIATE_COLLECTIONN},
	// 					{"localField", "uuid"},
	// 					{"foreignField", "campaignUUID"},
	// 					{"as", "candidates"},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$project", bson.D{
	// 					{"uuid", 0},
	// 					{"title", 0},
	// 					{"_id", 0},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$replaceRoot", bson.D{
	// 					{
	// 						"newRoot", bson.D{
	// 							{
	// 								"$mergeObject", bson.A{
	// 									bson.D{
	// 										{
	// 											"$arrayElemAt", bson.A{"$candidates", 0},
	// 										},
	// 									},
	// 									"$$ROOT",
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"project", bson.D{
	// 					{"candidateUUID", "uuid"},
	// 					{"uuid", 0},
	// 					{"title", 0},
	// 					{"_id", 0},
	// 				},
	// 			},
	// 		},
	// 	},
	// )

	if err != nil {

		panic(err)
	}

}
