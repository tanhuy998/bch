package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
	"app/infrastructure/restfull/common/crud"
)

type (
	delete struct {
		common.Controller
		crud.DeleteEndpointCurator
	}
)

func (this *delete) ANNOTATIONS_(
	auth.Authorize[constraint.TenantAgent],
) {

}
