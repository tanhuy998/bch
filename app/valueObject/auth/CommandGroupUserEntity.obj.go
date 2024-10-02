package authValueObject

import "app/model"

type (
	CommandGroupUserEntity struct {
		User  model.User
		Group model.CommandGroupUser
	}
)
