package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Tenant struct {
		UUID        uuid.UUID `json:"uuid" bson:"uuid"`
		Title       string    `json:"title" bson:"title"`
		Description string    `json:"description" bson:"description"`
	}

	TenantAgent struct {
		ProposedTime time.Time `json:"-" bson:"proposedTime"`
		Secret       []byte    `json:"-" bson:"secret"`
		UUID         uuid.UUID `json:"uuid" bson:"uuid"`
		TenantUUID   uuid.UUID `json:"tenantUUID" bson:"tenantUUID"`
		Username     string    `json:"username" bson:"username"`
		Password     string    `json:"password" bosn:"-'`
		Name         string    `json:"name" bson:"name"`
		Deactivated  bool      `json:"-" bson:"deactivated"`
	}
)

/*
IMPLEMENT passwordServiceAdapter.IPasswordDispatcher
*/
func (this *TenantAgent) GetRawUsername() []byte {

	return []byte(this.Username)
}
func (this *TenantAgent) GetRawPasword() []byte {

	return []byte(this.Password)
}

func (this *TenantAgent) GetSecret() []byte {

	return []byte(this.Secret)
}

func (this *TenantAgent) SetSecret(rawSecret []byte) {

	this.Secret = rawSecret
}

/*
	END IMPLEMENT passwordServiceAdapter.IPasswordDispatcher
*/
