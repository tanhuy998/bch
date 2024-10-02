package authServicePort

import (
	"app/model"
)

type (
	IGetCommandGroupUsers interface {
		Serve(groupUUID string) ([]*model.User, error)
	}
)
