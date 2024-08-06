package model

import "github.com/google/uuid"

type (
	User struct {
		UUID     uuid.UUID `json:"uuid" bson:"uuid"`
		UserName string    `json:"username" bson:"usename" validate:"required"`
		PassWord string    `json:"password" bson:"password" validate:"required"`
	}

	AuthGroup struct {
		UUID uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
		Name string    `json:"name" bson:"name" validate:"required"`
	}

	AuthGroupUser struct {
		UserUUID      uuid.UUID `json:"uuid" bson:"uuid"`
		AuthGroupUUID uuid.UUID `json:"authGroupUUID" bson:"authGroupUUID"`
	}
)
