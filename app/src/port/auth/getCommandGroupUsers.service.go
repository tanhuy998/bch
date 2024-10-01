package authServicePort

import (
	"app/src/model"
)

type (
	IGetCommandGroupUsers interface {
		Serve(groupUUID string) ([]*model.User, error)
	}
)
