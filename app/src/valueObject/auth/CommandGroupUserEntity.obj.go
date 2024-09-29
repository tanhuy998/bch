package authValueObject

import "app/domain/model"

type (
	CommandGroupUserEntity struct {
		User  model.User
		Group model.CommandGroupUser
	}
)
