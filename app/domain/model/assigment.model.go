package model

import "github.com/google/uuid"

type (
	AssigmentGroup struct {
		UUID uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
		Name string    `json:"name" bson:"name" validate:"requried"`
	}

	AssigmentGroupMember struct {
		UUID                 uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
		AssigmentGroupUUID   uuid.UUID `json:"assignmentGroupUUID" bson:"assignmentGroupUUID" validate:"required"`
		CommandGroupUserUUID uuid.UUID `json:"commandGroupUserUUID" bson:"commandGroupUserUUID" validate:"required"`
	}
)
