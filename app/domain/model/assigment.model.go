package model

import "github.com/google/uuid"

type (
	Assignment struct {
		UUID       uuid.UUID   `json:"uuid" bson:"uuid" validate:"required"`
		TenantUUID uuid.UUID   `json:"tenantUUID" bson:"tenantUUID" validate:"required"`
		OwnerShip  []uuid.UUID `json:"ownerShip" bson:"ownerShip"`
		Title      string      `json:"title" bson:"title"`
	}

	AssigmentGroup struct {
		UUID       uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
		TenantUUID uuid.UUID `json:"tenantUUID" bson:"tenantUUID" validate:"required"`
		Name       string    `json:"name" bson:"name" validate:"requried"`
	}

	AssigmentGroupMember struct {
		UUID                 uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
		TenantUUID           uuid.UUID `json:"tenantUUID" bson:"tenantUUID" validate:"required"`
		AssigmentGroupUUID   uuid.UUID `json:"assignmentGroupUUID" bson:"assignmentGroupUUID" validate:"required"`
		CommandGroupUserUUID uuid.UUID `json:"commandGroupUserUUID" bson:"commandGroupUserUUID" validate:"required"`
	}

	AssignmentTask struct {
		AssignmentGroupMemberUUID uuid.UUID   `json:"assignmentGroupMemberUUID" bson:"assigmnentGroupMemberUUID" validate:"required"`
		TenantUUID                uuid.UUID   `json:"tenantUUID" bson:"tenantUUID" validate:"required"`
		Payload                   interface{} `json:"payload" bson:"payload"`
	}
)
