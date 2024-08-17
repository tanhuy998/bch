package valueObject

import "app/domain/model"

type (
	ParticipatedCommandGroupDetail struct {
		CommandGroup *model.CommandGroup `json:"commandGroup" bson:"commandGroup"`
		Role         *model.Role         `json:"role" bson:"role"`
	}
)
