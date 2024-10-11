package valueObject

import "github.com/google/uuid"

type (
	ParticipatedCommandGroupReport struct {
		UserUUID   uuid.UUID                         `json:"userUUID" bson:"userUUID"`
		TenantUUID uuid.UUID                         `json:"tenantUUID" bson:"tenantUUID"`
		Details    []*ParticipatedCommandGroupDetail `json:"details" bson:"details"`
	}
)
