package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/crud"
)

type (
	delete struct {
		crud.DeleteEndpointCurator
		common.Controller
	}
)

func (this *delete) ENDPOINT_() {

}
