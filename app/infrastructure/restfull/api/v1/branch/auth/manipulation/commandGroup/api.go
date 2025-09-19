package commandGroup

import (
	"app/infrastructure/restfull/api/v1/branch/auth/manipulation/commandGroup/controller"
	"app/infrastructure/restfull/common"

	"github.com/kataras/iris/v12"
)

func API(parent iris.Party) {

	common.NewAPILauncher(parent).LaunchAPIOf(new(controller.CommandGroupController))
}
