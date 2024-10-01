package authValueObject

import "app/src/model"

type (
	CommandGroupUserEntity struct {
		User  model.User
		Group model.CommandGroupUser
	}
)
