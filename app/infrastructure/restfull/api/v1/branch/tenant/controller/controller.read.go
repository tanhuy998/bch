package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/restfull/common/crud"
)

type (
	read struct {
		crud.ReadEndpointCurator
		common.Controller
	}
)
