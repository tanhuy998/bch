package model

import "github.com/google/uuid"

type (
	CampaignAssignmentGroup struct {
		Name string     `json:"name" bson:"name" validate:"required"`
		UUID *uuid.UUID `json:"uuid" bson:"uuid"`
	}
)
