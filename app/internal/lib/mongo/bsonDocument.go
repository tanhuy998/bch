package libMongo

import (
	libCommon "app/internal/lib/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	BsonDocument struct {
		ObjectID *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	}
)

func (this *BsonDocument) GetObjectID() primitive.ObjectID {

	return *this.ObjectID
}

func (this *BsonDocument) SetObjectID(id primitive.ObjectID) {

	this.ObjectID = libCommon.PointerPrimitive(id)
}

func (this BsonDocument) GetCursor() interface{} /**tv(Cursor_T)*/ {

	return this.ObjectID
}
