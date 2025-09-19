package authSignaturesApi

import (
	"app/infrastructure/restfull/api/v1/branch/auth/signatures/controller"
	"app/infrastructure/restfull/common"

	"github.com/kataras/iris/v12/core/router"
)

func API(parent router.Party) {

	common.NewAPILauncher(parent).LaunchAPIOf(new(controller.SignatruresController))
}
