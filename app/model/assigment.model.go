package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Assignment struct {
		UUID       *uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		TenantUUID *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty"`
		CreatedAt  *time.Time `json:"createdAt" bson:"createdAt,omitempty"`
		CreatedBy  *uuid.UUID `json:"createdBy" bson:"createdBy,omitempty"`
		Deadline   *time.Time `json:"deadline" bson:"deadline,omitempty"`
		//OwnerShip  []uuid.UUID `json:"ownerShip" bson:"ownerShip"`
		Title      string `json:"title" bson:"title" validate:"required,omitempty"`
		Desciption string `json:"description" bson:"description,omitempty"`
	}

	AssignmentGroup struct {
		UUID             *uuid.UUID `json:"uuid,omitempty" bson:"uuid,omitempty"`
		AssignmentUUID   *uuid.UUID `json:"assignmentUUID" bson:"assignmentUUID"`
		TenantUUID       *uuid.UUID `json:"tenantUUID,omitempty" bson:"tenantUUID,omitempty"`
		CreatedBy        *uuid.UUID `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
		CommandGroupUUID *uuid.UUID `json:"commandGroupUUID,omitempty" bson:"commandGroupUUID,omitempty"`
		Name             string     `json:"name" bson:"name" validate:"required"`
	}

	AssignmentGroupMember struct {
		UUID                 *uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		TenantUUID           *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty"`
		AssigmentGroupUUID   *uuid.UUID `json:"assignmentGroupUUID" bson:"assignmentGroupUUID,omitempty"`
		CreatedBy            *uuid.UUID `json:"createdBy" bson:"createdBy"`
		CommandGroupUserUUID *uuid.UUID `json:"commandGroupUserUUID" bson:"commandGroupUserUUID,omitempty" validate:"required"`
	}

	AssignmentTask struct {
		AssignmentGroupMemberUUID *uuid.UUID `json:"assignmentGroupMemberUUID" bson:"assigmnentGroupMemberUUID" validate:"required"`
		TenantUUID                *uuid.UUID `json:"tenantUUID" bson:"tenantUUID" validate:"required"`
		Payload                   any        `json:"payload" bson:"payload"`
	}
)
