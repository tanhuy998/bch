package valueObject

import "github.com/google/uuid"

type (
	ParticipatedCommandGroupReport struct {
		UserUUID uuid.UUID                         `json:"userUUID"`
		Details  []*ParticipatedCommandGroupDetail `json:"details"`
	}
)
