package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Tenant struct {
		UUID        *uuid.UUID `json:"uuid" bson:"uuid"`
		Name        string     `bson:"title"`
		Description string     `bson:"description"`
	}

	TenantAgent struct {
		ProposedTime time.Time `json:"-" bson:"proposedTime,omitempty"`
		// Secret       []byte    `json:"-" bson:"secret"`
		UUID       *uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		TenantUUID *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty"`
		UserUUID   *uuid.UUID `json:"userUUID" bson:"userUUID,omitempty"`
		// Username     string    `json:"username,omitEmpty" bson:"username"`
		// Password     string    `json:"password,omitEmpty" bosn:"-'`
		// Name         string    `json:"name" bson:"name"`
		// Email        string    `json:"email" bson:"email"`
		Deactivated bool `json:"deactivated" bson:"deactivated"`
	}

	TenantAgentRegistration struct {
		UUID     uuid.UUID `bson:"uuid"`
		Username string    `json:"username" bson:"username"`
		Email    string    `json:"email" bson:"email"`
	}
)

// /*
// IMPLEMENT passwordServiceAdapter.IPasswordDispatcher
// */
// func (this *TenantAgent) GetRawUsername() []byte {

// 	return []byte(this.Username)
// }
// func (this *TenantAgent) GetRawPasword() []byte {

// 	return []byte(this.Password)
// }

// func (this *TenantAgent) GetSecret() []byte {

// 	return []byte(this.Secret)
// }

// func (this *TenantAgent) SetSecret(rawSecret []byte) {

// 	this.Secret = rawSecret
// }

/*
	END IMPLEMENT passwordServiceAdapter.IPasswordDispatcher
*/