package libMongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	BsonDocument struct {
		ObjectID *primitive.ObjectID `json:"_id" bson:"_id"`
	}
)

func (this BsonDocument) GetObjectID() primitive.ObjectID {

	return *this.ObjectID
}
