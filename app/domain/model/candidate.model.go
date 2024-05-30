package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	ObjectID     *primitive.ObjectID   `bson:"_id,omitempty"`
	UUID         *uuid.UUID            `json:"uuid" bson:"uuid,omitempty" validate:"required"`
	Name         *string               `json:"name" bson:"name,omitempty" validate:"required"`
	IDNumber     *string               `json:"idNumber" bson:"idNumber,omitempty" validate:"required, len=12"`
	Address      *string               `json:"address" bson:"uuid,omitempty" validate:"require"`
	SigningInfo  *CandidateSigningInfo `json:"signingInfo" bson:"signingInfo,omitempty"`
	CampaignUUID *uuid.UUID            `json:"campaignID" bson:"campaignUUID,omitempty"`
	Version      *time.Time            `json:"version" bson:"version,omitempty"`
}

func (this Candidate) GetObjectID() primitive.ObjectID {

	return *(this.ObjectID)
}
