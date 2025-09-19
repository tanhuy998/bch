package controller

import (
	"app/infrastructure/http/common"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
	"app/infrastructure/restfull/common/crud"
)

type (
	update struct {
		common.Controller
		crud.UpdateEndpointCurator
	}
)

func (this *update) ANNOTATIONS_(
	auth.Authorize[constraint.TenantAgent],
) {

}
