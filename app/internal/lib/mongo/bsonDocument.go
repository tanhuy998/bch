package libMongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	BsonDocument struct {
		ObjectID *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	}
)

func (this BsonDocument) GetObjectID() primitive.ObjectID {

	return *this.ObjectID
}
