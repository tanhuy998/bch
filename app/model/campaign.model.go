package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const MONGOD_CAMPAIGN_MODEL_COLLECTION = "campaigns"

type Campaign struct {
	IModel
	Model
	Title      string               `json:"title"`
	Time       time.Time            `json:"time"`
	ExpireAt   time.Time            `json:"expiredAt" bson:"expiredAt"`
	Candidates []primitive.ObjectID `json:"candidate_ids" bson:"candidate_ids"`
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
