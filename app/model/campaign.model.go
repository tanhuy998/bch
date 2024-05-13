package model

import (
	"time"

	"github.com/google/uuid"
)

const MONGOD_CAMPAIGN_MODEL_COLLECTION = "campaigns"

type Campaign struct {
	IModel
	Model
	UUID   uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
	Title  string    `json:"title" bson:"title" validate:"required"`
	Time   time.Time `json:"time" bson:"time"`
	Expire time.Time `json:"expire" bson:"expire" validate:"required"`
	//Candidates []primitive.ObjectID `json:"candidate_ids" bson:"candidate_ids"`
}

func (this *Campaign) CollectionName() string {

	return MONGOD_CAMPAIGN_MODEL_COLLECTION
}

func (this *Campaign) Store() {

	this.Establish()
}

func (this *Campaign) Insert() {

	this.Establish()
}

func (this *Campaign) Update() {

	this.Establish()
}

func (this *Campaign) Delete() {

	this.Establish()
}
