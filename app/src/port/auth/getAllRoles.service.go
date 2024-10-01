package authServicePort

import (
	"app/src/model"
)

type (
	IGetAllRoles interface {
		Serve() ([]*model.Role, error)
	}
)
