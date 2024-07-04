package model

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	ObjectID     *primitive.ObjectID   `bson:"_id,omitempty"`
	UUID         *uuid.UUID            `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Name         *string               `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	IDNumber     *string               `json:"idNumber,omitempty" bson:"idNumber,omitempty" validate:"required,number,len=12"`
	Address      *string               `json:"address,omitempty" bson:"address,omitempty" validate:"required"`
	Phone        *string               `json:"phone,omitempty" bson:"phone,omitempty"`
	SigningInfo  *CandidateSigningInfo `json:"signingInfo,omitempty" bson:"signingInfo,omitempty"`
	CampaignUUID *uuid.UUID            `json:"campaignUUID,omitempty" bson:"campaignUUID,omitempty"`
	Version      *time.Time            `json:"version,omitempty" bson:"version,omitempty"`
}

func (this Candidate) GetObjectID() primitive.ObjectID {

	return *(this.ObjectID)
}
