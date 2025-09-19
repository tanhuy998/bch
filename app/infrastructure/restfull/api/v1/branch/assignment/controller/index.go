package controller

import (
	"app/infrastructure/restfull/common"
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/middleware"
	"app/infrastructure/restfull/common/middleware/hook"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AssignmentController struct {
		common.CrudAPILauncher[*read, *create, *update, *delete]
	}
)

func (this *AssignmentController) AfterActivation(activator mvc.AfterActivation) {

	this.UseMiddleware(
		middleware.Auth(
			hook.AuthRequiredTenantAgentExceptOneOfRoles("COMMANDER"),
		),
	)
}

func (this *AssignmentController) ANNOTATIONS_(
	input.UseInputAuthorityMapping,
	input.UseInputTenantMapping,
) {
}
