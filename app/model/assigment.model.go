package model

import (
	libMongo "app/internal/lib/mongo"
	"time"

	"github.com/google/uuid"
)

type (
	Assignment struct {
		//TenantModel           `json:"-" bson:",inline"`
		Title                 string `json:"title" bson:"title" validate:"required"`
		Desciption            string `json:"description" bson:"description"`
		libMongo.BsonDocument `bson:",inline"`
		UUID                  *uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		CreatedAt             *time.Time `json:"createdAt" bson:"createdAt,omitempty"`
		Deadline              *time.Time `json:"deadline" bson:"deadline,omitempty"`
		CreatedBy             *uuid.UUID `json:"createdBy" bson:"createdBy,omitempty"`
		TenantUUID            *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty"`
		// join fields
		CreatedUser *User `json:"createdUser" bson:"createdUser,omitempty"`
		//OwnerShip  []uuid.UUID `json:"ownerShip" bson:"ownerShip"`
	}

	AssignmentGroup struct {
		Name             string     `json:"name" bson:"name" validate:"required"`
		UUID             *uuid.UUID `json:"uuid,omitempty" bson:"uuid,omitempty"`
		AssignmentUUID   *uuid.UUID `json:"assignmentUUID" bson:"assignmentUUID"`
		TenantUUID       *uuid.UUID `json:"tenantUUID,omitempty" bson:"tenantUUID,omitempty"`
		CreatedBy        *uuid.UUID `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
		CommandGroupUUID *uuid.UUID `json:"commandGroupUUID,omitempty" bson:"commandGroupUUID,omitempty"`
		// join fields
		CreatedUser  *User         `json:"createdUser" bson:"createdUser,omitempty"`
		CommandGroup *CommandGroup `json:"commandGroup" bson:"commandGroup,omitempty"`
	}

	AssignmentGroupMember struct {
		libMongo.BsonDocument `json:"-" bson:",inline"`
		UUID                  *uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		TenantUUID            *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty"`
		AssigmentGroupUUID    *uuid.UUID `json:"assignmentGroupUUID" bson:"assignmentGroupUUID,omitempty"`
		CreatedBy             *uuid.UUID `json:"createdBy" bson:"createdBy"`
		CommandGroupUserUUID  *uuid.UUID `json:"commandGroupUserUUID" bson:"commandGroupUserUUID,omitempty" validate:"required"`
		// join fields
		CommandGroupUser *CommandGroupUser `json:"commandGroupUser" bson:"commandGroupUser,omitempty"`
	}

	AssignmentTask struct {
		libMongo.BsonDocument     `json:"-" bson:",inline"`
		AssignmentGroupMemberUUID *uuid.UUID `json:"assignmentGroupMemberUUID" bson:"assigmnentGroupMemberUUID" validate:"required"`
		TenantUUID                *uuid.UUID `json:"tenantUUID" bson:"tenantUUID" validate:"required"`
		Payload                   any        `json:"payload" bson:"payload"`
	}
)
