package authServicePort

import (
	"app/model"
)

type (
	IGetAllRoles interface {
		Serve() ([]*model.Role, error)
	}
)
