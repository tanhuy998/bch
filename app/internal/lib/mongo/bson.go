package libMongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	IBsonDocument interface {
		GetObjectID() primitive.ObjectID
	}
)
