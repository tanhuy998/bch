package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/auth"
	"app/infrastructure/restfull/common/annotation/auth/constraint"
	"app/infrastructure/restfull/common/annotation/input"
)

type (
	UserManipualtionController struct {
		common.CrudAPILauncher[*read, *create, *update, *delete]
	}
)

func (this *UserManipualtionController) ANNOTATIONS_(
	auth.Authorize[constraint.TenantAgent],
	input.UseInputAuthorityMapping,
	input.UseInputTenantMapping,
) {

}
