package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/crud"
)

type (
	delete struct {
		common.Controller
		crud.DeleteEndpointCurator
	}
)
