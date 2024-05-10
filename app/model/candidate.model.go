package model

import "github.com/google/uuid"

type Candidate struct {
	IModel
	Model
	UUID        uuid.UUID            `json:"uuid" bson:"uuid" validate:"required"`
	IDNumber    string               `json:"idNumber" bson:"idNumber" validate:"required, len=12"`
	Address     string               `json:"address" bson:"uuid" validate:"require"`
	SigningInfo CandidateSigningInfo `json:"signingInfo" bson:"signingInfo"`
}
