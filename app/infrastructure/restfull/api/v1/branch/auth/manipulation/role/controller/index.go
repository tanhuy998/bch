package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
	"app/infrastructure/restfull/common/annotation/input"
)

type (
	AuthRoleManipulationController struct {
		common.CrudAPILauncher[*read, *create, *update, *delete]
	}
)

func (*AuthRoleManipulationController) ANNOTATIONS_(
	auth.Authorize[constraint.TenantAgent],
	input.UseInputAuthorityMapping,
	input.UseInputTenantMapping,
) {

}
