package model

import (
	"app/internal/generalToken"
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		*CommandGroup `bson:"groups,omitempty"`
		UUID          *uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		TenantUUID    *uuid.UUID `json:"tenantUUID" bson:"tenantUUID" validate:"required"`
		CreatedBy     *uuid.UUID `json:"createdBy" bson:"createdBy"`
		Name          string     `json:"name,omitempty" bson:"name"`
		Username      string     `json:"username" bson:"username" validate:"required"`
		PassWord      string     `json:"-" bson:"password" validate:"required"`
		Secret        []byte     `json:"-" bson:"secret"`
		IsDeactivated bool       `json:"deactivated" bson:"deactivated"`
		//OriginCommandGroupUUID *uuid.UUID `json:"-" bson:"originGroupUUIDCommandGroupUUID"`
		//Info          UserInfo  `json:"userInfo" bson:"userInfo"`
	}

	UserSession struct {
		UserUUID   *uuid.UUID                   `bson:"userUUID,omitempty"`
		TenantUUID *uuid.UUID                   `bson:"tenantUUID,omitempty"`
		SessionID  *generalToken.GeneralTokenID `bson:"sessionID,omitempty"`
		Expire     *time.Time                   `bson:"expire,omitempty"`
	}

	// UserInfo struct {
	// 	Name string `json:"name" bson:"name" validate:"required"`
	// }

	CommandGroup struct {
		UUID        *uuid.UUID `json:"uuid,omitempty" bson:"uuid,omitempty"`
		TenantUUID  *uuid.UUID `json:"tenantUUID" bson:"tenantUUID"`
		CreatedBy   *uuid.UUID `json:"createdBy" bson:"createdBy,omitempty"`
		Description string     `json:"description" bson:"description,omitempty"`
		Name        string     `json:"name" bson:"name,omitempty" validate:"required"`
	}

	CommandGroupUser struct {
		*User            `bson:"user,omitempty"`
		UUID             *uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		TenantUUID       *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty"`
		UserUUID         *uuid.UUID `json:"userUUID" bson:"userUUID,omitempty"`
		CommandGroupUUID *uuid.UUID `json:"commandGroupUUID" bson:"commandGroupUUID"`
		CreatedBy        *uuid.UUID `json:"createdBy" bson:"createdBy,omitempty"`
		Claims           []string   `json:"claims,omitempty" bson:"claims,omitempty"`
		// join fields
		CreatedUser *User `json:"createdUser" bson:"createdUser,omitempty"`
		//RoleUUID         uuid.UUID `json:"roleUUID" bson:"roleUUID"`
	}

	CommandGroupUserRole struct {
		*CommandGroupUser    `bson:"group,omitempty"`
		*Role                `bson:"roles,omitempty"`
		UUID                 *uuid.UUID `json:"uuid" bson:"uuid,omitempty" validate:"required"`
		TenantUUID           *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty" validate:"required"`
		CommandGroupUserUUID *uuid.UUID `json:"commandGroupUserUUID" bson:"commandGroupUserUUID,omitempty" validate:"required"`
		RoleUUID             *uuid.UUID `json:"roleUUID" bson:"roleUUID,omitempty" validate:"required"`
		CreatedBy            *uuid.UUID `json:"createdBy" bson:"createdBy,omitempty"`
	}

	Role struct {
		UUID       *uuid.UUID `json:"uuid,omitempty" bson:"uuid,omitempty"`
		TenantUUID *uuid.UUID `json:"tenantUUID,omitempty" bson:"tenantUUID,omitempty"`
		Name       string     `json:"name,omitempty" bson:"name,omitempty"`
	}
)

type (
	ParticipatedUsers struct {
	}
)

/*
IMPLEMENT passwordServiceAdapter.IPasswordDispatcher
*/
func (this *User) GetRawUsername() []byte {

	return []byte(this.Username)
}
func (this *User) GetRawPasword() []byte {

	return []byte(this.PassWord)
}

func (this *User) GetSecret() []byte {

	return []byte(this.Secret)
}

func (this *User) SetSecret(rawSecret []byte) {

	this.Secret = rawSecret
}

/*
	END IMPLEMENT passwordServiceAdapter.IPasswordDispatcher
*/
