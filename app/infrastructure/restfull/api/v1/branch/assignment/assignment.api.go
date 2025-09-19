package assignment

import (
	"app/infrastructure/restfull/api/v1/branch/assignment/controller"
	"app/infrastructure/restfull/common"

	"github.com/kataras/iris/v12"
)

func Legacy(parent iris.Party) {

	router := parent.Party("/assigns")

	launcher := common.NewAPILauncher(router)

	launcher.LaunchAPIOf(new(controller.AssignmentController))
}

func API(parent iris.Party) {

	launcher := common.NewAPILauncher(parent)

	launcher.LaunchAPIOf(new(controller.AssignmentController))
}
