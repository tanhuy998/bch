package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	IModel
	Model
	UUID        uuid.UUID            `json:"uuid" bson:"uuid" validate:"required"`
	Name        string               `json:"name" bson:"name" validate:"required"`
	IDNumber    string               `json:"idNumber" bson:"idNumber" validate:"required, len=12"`
	Address     string               `json:"address" bson:"uuid" validate:"require"`
	SigningInfo CandidateSigningInfo `json:"signingInfo" bson:"signingInfo"`
	CampaignID  primitive.ObjectID   `json:"campaignID" bson:"campaignID"`
}
