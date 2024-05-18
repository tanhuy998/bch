package model

import (
	"time"

	"github.com/google/uuid"
)

type Candidate struct {
	IModel
	Model
	UUID        uuid.UUID             `json:"uuid" bson:"uuid" validate:"required"`
	Name        *string               `json:"name" bson:"name" validate:"required"`
	IDNumber    *string               `json:"idNumber" bson:"idNumber" validate:"required, len=12"`
	Address     *string               `json:"address" bson:"uuid" validate:"require"`
	SigningInfo *CandidateSigningInfo `json:"signingInfo" bson:"signingInfo"`
	CampaignID  *uuid.UUID            `json:"campaignID" bson:"campaignID"`
	Version     *time.Time            `json:"version" bson:"version"`
}
