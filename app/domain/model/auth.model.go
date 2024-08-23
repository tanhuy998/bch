package model

import "github.com/google/uuid"

type (
	User struct {
		UUID          uuid.UUID `json:"uuid" bson:"uuid,omitempty"`
		Name          string    `json:"name,omitempty" bson:"name"`
		Username      string    `json:"-" bson:"username" validate:"required"`
		PassWord      string    `json:"-" bson:"password" validate:"required"`
		Secret        []byte    `json:"-" bson:"secret"`
		IsDeactivated bool      `json:"deactivated" bson:"deactivated"`
		//Info          UserInfo  `json:"userInfo" bson:"userInfo"`
	}

	// UserInfo struct {
	// 	Name string `json:"name" bson:"name" validate:"required"`
	// }

	CommandGroup struct {
		UUID uuid.UUID `json:"uuid,omitempty" bson:"uuid"`
		Name string    `json:"name" bson:"name" validate:"required"`
	}

	CommandGroupUser struct {
		UUID             uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
		UserUUID         uuid.UUID `json:"userUUID" bson:"userUUID"`
		CommandGroupUUID uuid.UUID `json:"commandGroupUUID" bson:"commandGroupUUID"`
		//RoleUUID         uuid.UUID `json:"roleUUID" bson:"roleUUID"`
	}

	CommandGroupUserRole struct {
		UUID                 uuid.UUID `json:"uuid" bson:"uuid" validate:"required"`
		CommandGroupUserUUID uuid.UUID `json:"commandGroupUserUUID" bson:"commandGroupUserUUID" validate:"required"`
		RoleUUID             uuid.UUID `json:"roleUUID" bson:"roleUUID" validate"required"`
	}

	Role struct {
		UUID uuid.UUID `json:"uuid" bson:"uuid"`
		Name string    `json:"name" bson:"name"`
	}
)
