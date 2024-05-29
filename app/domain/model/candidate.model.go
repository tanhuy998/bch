package model

import (
	"time"

	"github.com/google/uuid"
)

type Candidate struct {
	IModel
	Model
	UUID        *uuid.UUID            `json:"uuid" bson:"uuid,omitempty" validate:"required"`
	Name        *string               `json:"name" bson:"name,omitempty" validate:"required"`
	IDNumber    *string               `json:"idNumber" bson:"idNumber,omitempty" validate:"required, len=12"`
	Address     *string               `json:"address" bson:"uuid,omitempty" validate:"require"`
	SigningInfo *CandidateSigningInfo `json:"signingInfo" bson:"signingInfo,omitempty"`
	CampaignID  *uuid.UUID            `json:"campaignID" bson:"campaignID,omitempty"`
	Version     *time.Time            `json:"version" bson:"version,omitempty"`
}
