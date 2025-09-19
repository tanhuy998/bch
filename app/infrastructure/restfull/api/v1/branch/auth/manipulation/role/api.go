package role

import (
	"app/infrastructure/restfull/api/v1/branch/auth/manipulation/role/controller"
	"app/infrastructure/restfull/common"

	"github.com/kataras/iris/v12/core/router"
)

func API(parent router.Party) {

	common.NewAPILauncher(parent).LaunchAPIOf(new(controller.AuthRoleManipulationController))
}
