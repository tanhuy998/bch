package model

import (
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `json:"uuid" bson:"uuid"`
	UserName string    `json:"userName" bson:"userName"`
}
