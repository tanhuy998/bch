package valueObject

import (
	"app/model"

	"github.com/google/uuid"
)

type (
	ParticipatedCommandGroupDetail struct {
		CommandGroupUUID     uuid.UUID     `json:"commandGroupUUID" bson:"commandGroupUUID"`
		CommandGroupUserUUID uuid.UUID     `json:"commandGroupUserUUID" bson:"commandGroupUserUUID"`
		GroupName            string        `json:"groupName" bson:"name"`
		Roles                []*model.Role `json:"roles" bson:"roles"`
	}
)
