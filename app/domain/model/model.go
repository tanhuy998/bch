package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type IModel interface {
	GetObjectID() primitive.ObjectID
}
