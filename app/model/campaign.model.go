package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MONGOD_CAMPAIGN_MODEL_COLLECTION = "campaigns"

type Campaign struct {
	ObjectID  *primitive.ObjectID `bson:"_id,omitempty"`
	UUID      *uuid.UUID          `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Title     *string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	IssueTime *time.Time          `json:"issueTime,omitempty" bson:"issueTime,omitempty"`
	Expire    *time.Time          `json:"expire,omitempty" bson:"expire,omitempty" validate:"required"`
	Version   *time.Time          `json:"version,omitempty" bson:"version,omitempty"`
	//Candidates []primitive.ObjectID `json:"candidate_ids" bson:"candidate_ids"`
}

func (this Campaign) GetObjectID() primitive.ObjectID {

	return *(this.ObjectID)
}

// func (this *Campaign) CollectionName() string {

// 	return MONGOD_CAMPAIGN_MODEL_COLLECTION
// }

// func (this *Campaign) Store() {

// 	this.Establish()
// }

// func (this *Campaign) Insert() {

// 	this.Establish()
// }

// func (this *Campaign) Update() {

// 	this.Establish()
// }

// func (this *Campaign) Delete() {

// 	this.Establish()
// }
