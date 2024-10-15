package model

import (
	libMongo "app/internal/lib/mongo"

	"github.com/google/uuid"
)

type (
	IModel interface {
		libMongo.IBsonDocument
	}

	TenantDomainModel struct {
		TenantUUID *uuid.UUID `json:"tenantUUID" bson:"tenantUUID,omitempty"`
		CreatedBy  *uuid.UUID `json:"createdBy" bson:"createdBy,omitempty"`
	}
)
