package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/input"
)

type (
	CommandGroupController struct {
		common.CrudAPILauncher[*read, *create, *update, *delete]
	}
)

func (*CommandGroupController) ANNOTATIONS_(
	input.UseInputAuthorityMapping,
	input.UseInputTenantMapping,
) {

}
