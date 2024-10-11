package valueObject

import (
	"app/model"

	"github.com/google/uuid"
)

type (
	ParticipatedCommandGroupDetail struct {
		CommandGroupUUID uuid.UUID     `json:"commandGroupUUID" bson:"commandGroupUUID"`
		GroupName        string        `json:"groupName" bson:"name"`
		Roles            []*model.Role `json:"roles" bson:"roles"`
	}
)
